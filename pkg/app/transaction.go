package app

import (
	"context"

	"github.com/NpoolPlatform/sphinx-service/message/npool"
	"github.com/NpoolPlatform/sphinx-service/pkg/db"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

// 转账 / 提现
func ApplyTransaction(ctx context.Context, in *npool.ApplyTransactionRequest) (resp *npool.SuccessInfo, err error) {
	// preset
	resp = &npool.SuccessInfo{
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
	// rabbitmq notify
	// TODO
	return
}

// 交易状态查询
func GetInsiteTxStatus(ctx context.Context, in *npool.GetInsiteTxStatusRequest) (resp *npool.GetInsiteTxStatusResponse, err error) {
	return
}

// 余额查询
func GetBalance(ctx context.Context, in *npool.GetBalanceRequest) (resp *npool.AccountBalance, err error) {
	resp = &npool.AccountBalance{}
	return
}

// TODO: 账户交易查询
func GetTxJSON(ctx context.Context, in *npool.GetTxJSONRequest) (resp *npool.AccountTxJSON, err error) {
	return
}
