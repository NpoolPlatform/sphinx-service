package app

import (
	"context"

	trading "github.com/NpoolPlatform/message/npool/trading"
	"github.com/NpoolPlatform/sphinx-service/pkg/client"
	"github.com/NpoolPlatform/sphinx-service/pkg/db"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

// 转账 / 提现
func ApplyTransaction(ctx context.Context, in *trading.ApplyTransactionRequest) (resp *trading.SuccessInfo, err error) {
	// preset
	resp = &trading.SuccessInfo{
		Info: "aborted",
	}
	// check if transaction alreadey exists
	info, err := db.Client().Transaction.Query().
		Where(
			transaction.And(
				transaction.TransactionIDInsite(in.TransactionIdInsite),
			),
		).
		First(ctx)
	if info != nil {
		if info.AddressFrom == in.AddressFrom && info.AddressTo == info.AddressFrom && info.AmountUint64 == in.AmountUint64 {
			resp.Info = "success"
		} else {
			err = status.Error(codes.AlreadyExists, "transaction id insite already exists")
		}
		return
	}
	// check uuid signature
	if in.UuidSignature == "forbidden" {
		err = status.Error(codes.Canceled, "user signature invalid")
		return
	}
	// params check
	needManualReview := true
	txType := transaction.TypeUnknown
	if in.Type == "withdraw" {
		txType = transaction.TypeWithdraw
	} else if in.Type == "recharge" {
		txType = transaction.TypeRecharge
	} else if in.Type == "payment" {
		txType = transaction.TypePayment
	}
	// insert sql record
	info, err = db.Client().Transaction.Create().
		SetCoinID(in.CoinId).
		SetMutex(false).
		SetStatus(transaction.StatusPendingReview).
		SetTransactionIDChain(in.TransactionIdInsite).
		SetAmountUint64(in.AmountUint64).
		SetAmountFloat64(in.AmountFloat64).
		SetAddressFrom(in.AddressFrom).
		SetAddressTo(in.AddressTo).
		SetNeedManualReview(needManualReview).
		SetSignatureUser(in.UuidSignature).
		SetType(txType).
		Save(ctx)
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
	err = client.LetApproveTransaction(info.TransactionIDInsite)
	if err != nil {
		err = status.Error(codes.Internal, "cannot notify transaction approval service")
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
	return
}

// 余额查询
func GetBalance(ctx context.Context, in *trading.GetBalanceRequest) (resp *trading.AccountBalance, err error) {
	resp = &trading.AccountBalance{}
	return
}

// TODO: 账户交易查询
func GetTxJSON(ctx context.Context, in *trading.GetTxJSONRequest) (resp *trading.AccountTxJSON, err error) {
	return
}
