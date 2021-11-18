package app

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/trading"
	"github.com/NpoolPlatform/sphinx-service/pkg/client"
	"github.com/NpoolPlatform/sphinx-service/pkg/crud"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

// 创建账号
func RegisterAccount(ctx context.Context, coinTypeID int32, uuid string) (account *trading.AccountAddress, err error) {
	// online mode
	// coinType := sphinxplugin.CoinType(coinTypeID)
	// resp, err := client.ClientProxy.WalletNew(ctx, &signproxy.WalletNewRequest{
	// 	UUID: uuid,
	// 	CoinType: coinType,
	// })
	// "github.com/NpoolPlatform/message/npool/signproxy"
	// debug mode
	resp, err := client.ClientProxy.WalletNew(ctx, &emptypb.Empty{})
	if err != nil {
		resp = nil
		return
	}
	account = &trading.AccountAddress{
		CoinId:  coinTypeID,
		Address: resp.Info.Address,
		Uuid:    uuid,
	}
	return
}

// 余额查询
func GetBalance(ctx context.Context, in *trading.GetBalanceRequest) (resp *trading.AccountBalance, err error) {
	respRPC, err := client.ClientProxy.WalletBalance(ctx, &sphinxplugin.WalletBalanceRequest{
		CoinType: sphinxplugin.CoinType(in.CoinId),
		Address:  in.Address,
	})
	if err != nil {
		err = status.Errorf(codes.Internal, "get wallet balance failed, %v", err)
		return
	}
	amountUint64, amountString, amountFloat64 := UntestedDecomposeStringAmount(respRPC.Info.Balance)
	logger.Sugar().Infof("amount in: %v", respRPC.Info.Balance)
	logger.Sugar().Infof("amount decomposed: %v", amountString)
	resp = &trading.AccountBalance{
		CoinId:        in.CoinId,
		Address:       in.Address,
		TimestampUtc:  time.Now().UTC().Unix(),
		AmountFloat64: amountFloat64,
		AmountUint64:  amountUint64,
	}
	return
}

// 转账 / 提现
func ApplyTransaction(ctx context.Context, in *trading.ApplyTransactionRequest) (resp *trading.SuccessInfo, err error) {
	// preset
	resp = &trading.SuccessInfo{
		Info: "aborted",
	}
	isExisted, err := crud.CheckRecordIfExistTransaction(in)
	if err != nil {
		return
	} else if isExisted {
		resp.Info = "success"
		return
	}
	// check uuid signature
	if in.UuidSignature == "forbidden" {
		err = status.Error(codes.Canceled, "user signature invalid")
		return
	}
	// if amount > xxx, needManualReview => true, and etc.
	needManualReview := true
	// convert type
	txType := transaction.TypeUnknown
	if in.Type == "withdraw" {
		txType = transaction.TypeWithdraw
	} else if in.Type == "recharge" {
		txType = transaction.TypeRecharge
	} else if in.Type == "payment" {
		txType = transaction.TypePayment
	}
	// insert sql record
	info, err := crud.CreateRecordTransaction(in, needManualReview, txType)
	if err != nil {
		err = status.Error(codes.Internal, "database error")
		return
	}
	/*
		// rabbitmq notify
		// err = server.PublishNotificationTransactionCreate(&message.NotificationTransaction{
		// 	ID:                  info.ID,
		// 	TransactionIDInsite: info.TransactionIDInsite,
		// })
		// if err != nil {
		// 	err = status.Errorf(codes.Internal, "publishing notification error, check rabbitmq: %v", err)
		// } else {
		// 	resp.Info = "success"
		// }
	*/
	// grpc call approval service
	err = LetApproveTransaction(info)
	if err != nil {
		err = status.Errorf(codes.Internal, "cannot notify transaction approval service, error: %v", err)
		return
	}
	// done
	return resp, err
}

// TODO: 账户交易查询
func GetTxJSON(ctx context.Context, in *trading.GetTxJSONRequest) (resp *trading.AccountTxJSON, err error) {
	return
}

// 交易状态查询
func GetInsiteTxStatus(ctx context.Context, in *trading.GetInsiteTxStatusRequest) (resp *trading.GetInsiteTxStatusResponse, err error) {
	// MARK 交给钱包代理
	return
}
