package crud

import (
	"github.com/NpoolPlatform/message/npool/trading" //nolint
	"github.com/NpoolPlatform/sphinx-service/pkg/db"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func CreateRecordTransaction(in *trading.ApplyTransactionRequest, needManualReview bool, txType transaction.Type) (info *ent.Transaction, err error) {
	tmpCoinInfo, err := db.Client().CoinInfo.Query().Where(coininfo.Name(in.CoinName)).First(ctxPublic)
	if err != nil {
		info = nil
		return
	}
	info, err = db.Client().Transaction.Create().
		SetCoin(tmpCoinInfo).
		SetMutex(false).
		SetStatus(transaction.StatusPendingReview).
		SetTransactionIDChain(in.TransactionIdInsite).
		SetAmountFloat64(in.AmountFloat64).
		SetAddressFrom(in.AddressFrom).
		SetAddressTo(in.AddressTo).
		SetNeedManualReview(needManualReview).
		SetSignatureUser(in.UuidSignature).
		SetType(txType).
		Save(ctxPublic)
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
		First(ctxPublic)
	if info != nil { // has record
		isExisted = true
		if info.AddressFrom != in.AddressFrom || info.AddressTo != info.AddressFrom {
			err = status.Error(codes.AlreadyExists, "transaction id insite already exists")
		}
	}
	return
}
