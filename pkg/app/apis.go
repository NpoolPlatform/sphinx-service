package app

import (
	"context"
	"time"

	"github.com/NpoolPlatform/message/npool/trading"
	"github.com/NpoolPlatform/sphinx-service/pkg/crud"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

// 创建账号
func RegisterAccount(ctx context.Context, coinName, uuid string) (account *trading.AccountAddress, err error) {
	account = nil
	_, err = crud.CoinName2CoinID(coinName)
	if err != nil {
		return
	}
	address, err := LetCreateAccount(coinName, uuid)
	if err != nil {
		return
	}
	account = &trading.AccountAddress{
		CoinName: coinName,
		Address:  address,
		Uuid:     uuid,
	}
	return
}

// 余额查询
func GetBalance(ctx context.Context, in *trading.GetBalanceRequest) (resp *trading.AccountBalance, err error) {
	balance, err := LetGetWalletBalance(in.CoinName, in.Address)
	if err != nil {
		return
	}
	resp = &trading.AccountBalance{
		CoinName:      in.CoinName,
		Address:       in.Address,
		TimestampUtc:  time.Now().UTC().Unix(),
		AmountFloat64: balance,
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
