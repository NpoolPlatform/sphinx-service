package crud

import (
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/price"
	"github.com/NpoolPlatform/message/npool/signproxy"
	"github.com/NpoolPlatform/message/npool/trading" //nolint
	"github.com/NpoolPlatform/sphinx-service/pkg/db"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

func CreateRecordTransaction(in *trading.CreateTransactionRequest, needManualReview bool, txType transaction.Type) (info *ent.Transaction, err error) {
	tmpCoinInfo, err := db.Client().CoinInfo.Query().Where(coininfo.Name(in.CoinName)).Only(ctxPublic)
	if err != nil {
		logger.Sugar().Warn(in.CoinName, "coin not found", err)
		return
	}
	info, err = db.Client().Transaction.Create().
		SetAmountUint64(price.VisualPriceToDBPrice(in.AmountFloat64)).
		SetAmountFloat64(in.AmountFloat64).
		SetAddressFrom(in.AddressFrom).
		SetAddressTo(in.AddressTo).
		SetNeedManualReview(needManualReview).
		SetType(txType).
		SetTransactionIDInsite(in.TransactionIdInsite).
		SetTransactionIDChain("").
		SetStatus(transaction.StatusPendingReview).
		SetMutex(false).
		SetSignatureUser(in.UuidSignature).
		SetSignaturePlatform("test-version-direct-pass").
		SetCreatetimeUtc(time.Now().UTC().Unix()).
		SetUpdatetimeUtc(time.Now().UTC().Unix()).
		SetCoin(tmpCoinInfo).
		Save(ctxPublic)
	return
}

func CheckRecordIfExistTransaction(in *trading.CreateTransactionRequest) (isExisted bool, err error) {
	var info []*ent.Transaction
	info, err = db.Client().Transaction.Query().
		Where(
			transaction.And(
				transaction.TransactionIDInsite(in.TransactionIdInsite),
			),
		).All(ctxPublic)
	if len(info) > 0 { // has record, definitely len == 1
		isExisted = true
		if info[0].AddressFrom != in.AddressFrom || info[0].AddressTo != info[0].AddressFrom {
			err = status.Error(codes.AlreadyExists, "transaction id insite already exists")
		}
	}
	return
}

func UpdateTransactionStatus(in *trading.ACKRequest) (isSuccess bool, err error) {
	isSuccess = true
	entResp, err := db.Client().Transaction.Query().
		Where(
			transaction.And(
				transaction.TransactionIDInsite(in.TransactionIdInsite),
			),
		).
		First(ctxPublic)
	if err != nil || entResp == nil {
		logger.Sugar().Errorf("transaction incorrect, %v", err)
		return
	}
	flagErr := 0
	if in.TransactionType == signproxy.TransactionType_TransactionNew {
		if entResp.Status != transaction.StatusPendingProcess {
			flagErr = 1
		}
		err = entResp.Update().
			SetStatus(transaction.StatusPendingSigninfo).
			SetMutex(false).
			Exec(ctxPublic)
	} else if in.TransactionType == signproxy.TransactionType_PreSign {
		if entResp.Status != transaction.StatusPendingSigninfo {
			flagErr = 1
		}
		err = entResp.Update().
			SetStatus(transaction.StatusPendingSign).
			SetMutex(false).
			Exec(ctxPublic)
	} else if in.TransactionType == signproxy.TransactionType_Signature {
		if entResp.Status != transaction.StatusPendingSign {
			flagErr = 1
		}
		err = entResp.Update().
			SetStatus(transaction.StatusPendingBroadcast).
			SetMutex(false).
			Exec(ctxPublic)
	} else if in.TransactionType == signproxy.TransactionType_Broadcast {
		if entResp.Status != transaction.StatusPendingBroadcast {
			flagErr = 1
		}
		err = entResp.Update().
			SetStatus(transaction.StatusPendingConfirm).
			SetTransactionIDChain(in.TransactionIdChain).
			SetMutex(false).
			Exec(ctxPublic)
	}
	if flagErr == 1 {
		logger.Sugar().Errorf("failed to update transaction status, %v", err)
		isSuccess = false
	}
	if err != nil {
		logger.Sugar().Errorf("update transaction db failed: %v", err)
		isSuccess = false
	}
	return isSuccess, err
}
