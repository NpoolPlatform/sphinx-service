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

func CreateWallet(ctx context.Context, coinName, uuid string) (resp *trading.CreateWalletResponse, err error) {
	// Check if coin exists
	coinInfo, err := crud.CoinName2Coin(ctx, coinName)
	if err != nil {
		return
	}

	// Gen one-time ID
	tmpTxID := rules.GenerateTID4CreateWallet(coinInfo, uuid)

	// Push through rabbitmq to signproxy
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

	// Wait for signproxy reply (or auto-timeout)
	ackResp, err := ListenTillSucceeded(tmpTxID)
	if err != nil {
		err = xerrors.Errorf("listener error: %v", err)
		return
	}

	// Struct return
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

func GetWalletBalance(ctx context.Context, in *trading.GetWalletBalanceRequest) (resp *trading.GetWalletBalanceResponse, err error) {
	// Check coin
	coinInfo, err := crud.CoinName2Coin(ctx, in.Info.CoinName)
	if err != nil {
		return
	}

	// Gen one-time id
	tmpTxID := rules.GenerateTID4GetWalletBalance(coinInfo, in.Info.Address)

	// Push through rabbitmq to signproxy
	err = server.PublishDefaultNotification(&message.NotificationTransaction{
		CoinType:            sphinxplugin.CoinType(coinInfo.CoinTypeID),
		TransactionType:     signproxy.TransactionType_Balance,
		TransactionIDInsite: tmpTxID,
		AmountFloat64:       0, // no need
		AddressFrom:         in.Info.Address,
		AddressTo:           in.Info.Address,
		TransactionIDChain:  "",                      // no need
		SignatureUser:       "",                      // no need
		SignaturePlatform:   "",                      // no need
		CreatetimeUtc:       time.Now().UTC().Unix(), // no need
		UpdatetimeUtc:       time.Now().UTC().Unix(), // no need
		IsSuccess:           false,                   // no need
		IsFailed:            false,                   // no need
	})
	if err != nil {
		err = xerrors.Errorf("message publish error: %v", err)
		return
	}

	// Wait for signproxy reply
	ackResp, err := ListenTillSucceeded(tmpTxID)
	if err != nil {
		err = xerrors.Errorf("listener error: %v", err)
		return
	}

	// Struct return
	resp = &trading.GetWalletBalanceResponse{
		Info: &trading.EntAccount{
			CoinName: in.Info.CoinName,
			Address:  in.Info.Address,
		},
		AmountFloat64: ackResp.Balance,
	}
	return resp, err
}

func CreateTransaction(ctx context.Context, in *trading.CreateTransactionRequest) (resp *trading.CreateTransactionResponse, err error) {
	// Check uuid signature next-version
	if in.UUIDSignature == "forbidden" {
		err = status.Error(codes.Canceled, "user signature invalid")
		return
	}

	// Mocked auto-review logic
	/*
		next-version: set this field by amount, login_ip, etc.
	*/
	needManualReview := true

	// Convert type
	txType := transaction.TypeUnknown
	if in.Info.InsiteTxType == "withdraw" {
		txType = transaction.TypeWithdraw
	} else if in.Info.InsiteTxType == "recharge" {
		txType = transaction.TypeRecharge
	} else if in.Info.InsiteTxType == "payment" {
		txType = transaction.TypePayment
	}

	// Insert sql record
	info, err := crud.CreateRecordTransaction(in, needManualReview, txType)
	if err != nil {
		// if same transaction(and uuid), err == nil, return old record.
		// if TID exists but not same, err happens here;
		return
	}

	// Grpc call approval service
	err = MockApproveTransaction(info)
	if err != nil {
		err = status.Errorf(codes.Internal, "cannot notify transaction approval service, error: %v", err)
		return
	}
	return resp, err
}

func GetTransaction(ctx context.Context, in *trading.GetTransactionRequest) (resp *trading.GetTransactionResponse, err error) {
	transactionRow, err := crud.GetTransaction(ctx, in)
	if err == nil {
		// Judge success/fail from transaction current status
		var flagFailed, flagSucceeded bool
		if transactionRow.Status == transaction.StatusDone {
			flagSucceeded = true
		} else if transactionRow.Status == transaction.StatusRejected ||
			transactionRow.Status == transaction.StatusError ||
			transactionRow.Status == transaction.StatusErrorExpected {
			flagFailed = true
		}
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
			Succeeded:          flagSucceeded,
			Failed:             flagFailed,
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
