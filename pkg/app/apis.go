package app

import (
	"context"
	"time"

	cv "github.com/NpoolPlatform/go-service-framework/pkg/version"
	"github.com/NpoolPlatform/message/npool/signproxy"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/trading" //nolint
	"github.com/NpoolPlatform/sphinx-service/pkg/crud"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	"github.com/NpoolPlatform/sphinx-service/pkg/message/message"
	"github.com/NpoolPlatform/sphinx-service/pkg/message/server"
	"github.com/NpoolPlatform/sphinx-service/pkg/rules"
	"github.com/gogo/status"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
)

// 创建账号 done
func CreateWallet(ctx context.Context, coinName, uuid string) (resp *trading.CreateWalletResponse, err error) {
	// check coin
	coinInfo, err := crud.CoinName2Coin(ctx, coinName)
	if err != nil {
		return
	}
	// gen one-time ID
	tmpTxID := rules.GenerateTID4CreateWallet(coinInfo, uuid)
	// push through rabbitmq to signproxy
	err = server.PublishDefaultNotification(&message.NotificationTransaction{
		CoinType:            sphinxplugin.CoinType(coinInfo.CoinTypeID),
		TransactionType:     signproxy.TransactionType_WalletNew,
		TransactionIDInsite: tmpTxID,
		AmountFloat64:       0,                       // no need
		AddressFrom:         "",                      // no need
		AddressTo:           "",                      // no need
		TransactionIDChain:  "",                      // no need
		SignatureUser:       "",                      // no need
		SignaturePlatform:   "",                      // next-version
		CreatetimeUtc:       time.Now().UTC().Unix(), // no need
		UpdatetimeUtc:       time.Now().UTC().Unix(), // no need
		IsSuccess:           false,                   // no need
		IsFailed:            false,                   // no need
	})
	if err != nil {
		err = xerrors.Errorf("message publish error: %v", err)
		return
	}
	// wait for signproxy reply
	ackResp, err := ListenTillSucceeded(tmpTxID)
	if err != nil {
		err = xerrors.Errorf("listener error: %v", err)
		return
	}
	// struct return
	if ackResp.Address == "" {
		err = xerrors.New("empty reply from server when creating account")
	} else {
		resp = &trading.CreateWalletResponse{
			Info: &trading.EntAccount{
				CoinName: coinName,
				Address:  ackResp.Address,
			},
		}
	}
	return resp, err
}

// 余额查询 done
func GetWalletBalance(ctx context.Context, in *trading.GetWalletBalanceRequest) (resp *trading.GetWalletBalanceResponse, err error) {
	// check coin
	coinInfo, err := crud.CoinName2Coin(ctx, in.Info.CoinName)
	if err != nil {
		return
	}
	// gen one-time id
	tmpTxID := rules.GenerateTID4GetWalletBalance(coinInfo, in.Info.Address)
	// push through rabbitmq to signproxy
	err = server.PublishDefaultNotification(&message.NotificationTransaction{
		CoinType:            sphinxplugin.CoinType(coinInfo.CoinTypeID),
		TransactionType:     signproxy.TransactionType_Balance,
		TransactionIDInsite: tmpTxID,
		AddressFrom:         in.Info.Address,
		AddressTo:           in.Info.Address,
		CreatetimeUtc:       time.Now().UTC().Unix(),
		UpdatetimeUtc:       time.Now().UTC().Unix(),
	})
	if err != nil {
		err = xerrors.Errorf("message publish error: %v", err)
		return
	}
	// wait for signproxy reply
	ackResp, err := ListenTillSucceeded(tmpTxID)
	if err != nil {
		err = xerrors.Errorf("listener error: %v", err)
		return
	}
	// struct return
	resp = &trading.GetWalletBalanceResponse{
		Info: &trading.EntAccount{
			CoinName: in.Info.CoinName,
			Address:  in.Info.Address,
		},
		AmountFloat64: ackResp.Balance,
	}
	return resp, err
}

// 转账 / 提现 pending
func CreateTransaction(ctx context.Context, in *trading.CreateTransactionRequest) (resp *trading.CreateTransactionResponse, err error) {
	// preset
	resp = &trading.CreateTransactionResponse{
		Info: in.Info,
	}
	_, err = crud.CheckRecordIfExistTransaction(in)
	if err != nil {
		return
	}
	// check uuid signature
	if in.UUIDSignature == "forbidden" {
		err = status.Error(codes.Canceled, "user signature invalid")
		return
	}
	// if amount > xxx, needManualReview => true, and etc.
	needManualReview := true
	// convert type
	txType := transaction.TypeUnknown
	if in.Info.InsiteTxType == "withdraw" {
		txType = transaction.TypeWithdraw
	} else if in.Info.InsiteTxType == "recharge" {
		txType = transaction.TypeRecharge
	} else if in.Info.InsiteTxType == "payment" {
		txType = transaction.TypePayment
	}
	// insert sql record
	info, err := crud.CreateRecordTransaction(in, needManualReview, txType)
	if err != nil {
		return
	}
	// grpc call approval service
	err = MockApproveTransaction(info)
	if err != nil {
		err = status.Errorf(codes.Internal, "cannot notify transaction approval service, error: %v", err)
		return
	}
	// done
	return resp, err
}

// 交易状态查询 next-version done for now
func GetTransaction(ctx context.Context, in *trading.GetTransactionRequest) (resp *trading.GetTransactionResponse, err error) {
	transactionRow, err := crud.GetTransaction(ctx, in)
	// TODO: Implement succeeded and failed judgement
	if err == nil {
		coinName := transactionRow.Edges.Coin
		resp = &trading.GetTransactionResponse{
			Info: &trading.BaseTx{
				TransactionIDInsite: transactionRow.TransactionIDInsite,
				CoinName:            coinName.Name,
				AmountFloat64:       transactionRow.AmountFloat64,
				AddressFrom:         transactionRow.AddressFrom,
				AddressTo:           transactionRow.AddressTo,
				InsiteTxType:        transactionRow.Type.String(),
				CreatetimeUTC:       transactionRow.CreatetimeUtc,
			},
			UpdatetimeUTC:      transactionRow.UpdatetimeUtc,
			Succeeded:          false,
			Failed:             false,
			TransactionIDChain: transactionRow.TransactionIDChain,
			Status:             string(transactionRow.Status),
		}
	}
	return
}

// Version (original code)
func Version() (*trading.VersionResponse, error) {
	info, err := cv.GetVersion()
	if err != nil {
		return nil, xerrors.Errorf("get service version error: %w", err)
	}
	return &trading.VersionResponse{
		Info: info,
	}, nil
}
