package crud

import (
	"context"

	"github.com/NpoolPlatform/message/npool/trading"
	"github.com/NpoolPlatform/sphinx-service/pkg/db"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

var ctx context.Context

func init() {
	ctx = context.Background()
}

func CreateRecordTransaction(in *trading.ApplyTransactionRequest, needManualReview bool, txType transaction.Type) (info *ent.Transaction, err error) {
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
	return
}

func CheckRecordIfExistTransaction(in *trading.ApplyTransactionRequest) (isExisted bool, err error) {
	var info *ent.Transaction
	info, err = db.Client().Transaction.Query().
		Where(
			transaction.And(
				transaction.TransactionIDInsite(in.TransactionIdInsite),
			),
		).
		First(ctx)
	if info != nil { // has record
		isExisted = true
		if info.AddressFrom != in.AddressFrom || info.AddressTo != info.AddressFrom || info.AmountUint64 != in.AmountUint64 {
			err = status.Error(codes.AlreadyExists, "transaction id insite already exists")
		}
	}
	return
}
