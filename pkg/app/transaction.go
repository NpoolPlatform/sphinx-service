package app

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/price"
	"github.com/NpoolPlatform/message/npool/signproxy"
	"github.com/NpoolPlatform/message/npool/trading"
	"github.com/NpoolPlatform/sphinx-service/pkg/client"
	"github.com/NpoolPlatform/sphinx-service/pkg/crud"
	"github.com/NpoolPlatform/sphinx-service/pkg/db"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	"github.com/NpoolPlatform/sphinx-service/pkg/message/message"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

const priceScale = 1000000000000

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
		err = status.Errorf(codes.Internal, "cannot notify transaction approval service", err)
		return
	}
	// MARK: approve result override
	_, err = db.Client().Transaction.Update().
		SetStatus(transaction.StatusPendingReview).
		SetMutex(false).
		Save(ctx)
	if err != nil {
		err = status.Error(codes.Internal, "database error")
		return
	}
	// done
	return resp, err
}

// 交易状态查询
func GetInsiteTxStatus(ctx context.Context, in *trading.GetInsiteTxStatusRequest) (resp *trading.GetInsiteTxStatusResponse, err error) {
	// MARK 交给钱包代理
	return
}

// 余额查询
func GetBalance(ctx context.Context, in *trading.GetBalanceRequest) (resp *trading.AccountBalance, err error) {
	client.ClientProxy.WalletBalance(ctx, &signproxy.WalletBalanceRequest{
		CoinType: 0,
		Address:  in.Address,
	})
	respRPC, err := client.ClientProxy.WalletBalance(ctx, &signproxy.WalletBalanceRequest{
		CoinType: signproxy.CoinType(in.CoinId),
		Address:  in.Address,
	})
	if err != nil {
		err = status.Errorf(codes.Internal, "get wallet balance failed, %w", err)
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

// 金额转换函数
func UntestedDecomposeStringAmount(str string) (amountUint64 uint64, amountString string, amountFloat64 float64) {
	// get value
	// for initial result: str == target_x_fil*10^18
	bi, _ := new(big.Int).SetString(str, 10)
	// prepare for division
	filExp := new(big.Int)
	filExp.Exp(big.NewInt(10), big.NewInt(9), nil)
	// make a float copy
	bf := new(big.Float).SetInt(bi)
	// divide
	bf.Quo(bf, new(big.Float).SetInt(filExp))
	amountString = fmt.Sprintf("%f", bf)
	amountFloat64, _ = bf.Float64()
	amountUint64 = price.VisualPriceToDBPrice(amountFloat64)
	return
}

// TODO: 账户交易查询
func GetTxJSON(ctx context.Context, in *trading.GetTxJSONRequest) (resp *trading.AccountTxJSON, err error) {
	return
}

// 审核服务 进行审核 overrided
func LetApproveTransaction(tx *ent.Transaction) (err error) {
	// MARK: currently overrided by pkg/app/transaction.go
	return
}

// 钱包代理 查询交易状态
func CheckIfTransactionComplete(tx *ent.Transaction) (err error) {
	return
}

// 钱包代理 创建账号
func RegisterAccount(coinTypeID int32, uuid string) (account *trading.AccountAddress, err error) {
	coinType := signproxy.CoinType(coinTypeID)
	notification := &message.NotificationTransaction{
		TransactionType: signproxy.TransactionType_TransactionTypeCreateAccount,
		CoinType:        coinType,
		UUID:            uuid,
		CreatetimeUtc:   time.Now().UTC().Unix(),
		UpdatetimeUtc:   0,
		IsSuccess:       false,
		IsFailed:        false,
	}
	return
}

// 钱包代理 进行转账

// 钱包代理 查询余额
