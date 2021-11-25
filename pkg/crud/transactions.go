package crud

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/price"
	"github.com/NpoolPlatform/message/npool/signproxy"
	"github.com/NpoolPlatform/message/npool/trading" //nolint
	"github.com/NpoolPlatform/sphinx-service/pkg/db"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/coininfo"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	"golang.org/x/xerrors"
)

var ctxPublic context.Context

func init() {
	ctxPublic = context.Background()
}

func CreateRecordTransaction(in *trading.CreateTransactionRequest, needManualReview bool, txType transaction.Type) (info *ent.Transaction, err error) {
	// get coin info
	coinInfo, err := db.Client().CoinInfo.Query().Where(coininfo.Name(in.Info.CoinName)).Only(ctxPublic)
	if err != nil {
		err = xerrors.Errorf(in.Info.CoinName+" coin not found %v", err)
		return
	}
	// check if exists
	info, err = GetTransactionOrNil(in)
	if info != nil {
		err = xerrors.Errorf("tx already exists: %v", info)
		return
	} else if err != nil {
		err = xerrors.Errorf("db error: %v", err)
		return
	}
	// do create
	info, err = db.Client().Transaction.Create().
		SetAmountUint64(price.VisualPriceToDBPrice(in.Info.AmountFloat64)).
		SetAmountFloat64(in.Info.AmountFloat64).
		SetAddressFrom(in.Info.AddressFrom).
		SetAddressTo(in.Info.AddressTo).
		SetNeedManualReview(needManualReview).
		SetType(txType).
		SetTransactionIDInsite(in.Info.TransactionIDInsite).
		SetTransactionIDChain("").
		SetStatus(transaction.StatusPendingReview).
		SetMutex(false).
		SetSignatureUser(in.UUIDSignature).
		SetSignaturePlatform("test-version-direct-pass").
		SetCreatetimeUtc(time.Now().UTC().Unix()).
		SetUpdatetimeUtc(time.Now().UTC().Unix()).
		SetCoin(coinInfo).
		Save(ctxPublic)
	return info, err
}

func GetTransactionOrNil(in *trading.CreateTransactionRequest) (record *ent.Transaction, err error) {
	var info []*ent.Transaction
	info, err = db.Client().Transaction.Query().
		Where(
			transaction.And(
				transaction.TransactionIDInsite(in.Info.TransactionIDInsite),
			),
		).All(ctxPublic)
	if err != nil {
		return
	}
	if len(info) > 0 {
		if len(info) > 1 {
			err = xerrors.New("impossible")
		}
		coinInfo, err := info[0].QueryCoin().Only(ctxPublic)
		if err != nil {
			err = xerrors.Errorf("transaction no coininfo %v", err)
			return record, err
		}
		if info[0].AddressFrom == in.Info.AddressFrom &&
			info[0].AddressTo == in.Info.AddressTo &&
			info[0].AmountFloat64 == in.Info.AmountFloat64 &&
			coinInfo.Name == in.Info.CoinName {
			record = info[0]
		}
	}
	return record, err
}

func UpdateTransactionStatus(ctx context.Context, in *trading.ACKRequest) (isSuccess bool, err error) {
	isSuccess = true
	entResp, err := db.Client().Transaction.Query().
		Where(
			transaction.And(
				transaction.TransactionIDInsite(in.TransactionIdInsite),
			),
		).
		First(ctx)
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
			Exec(ctx)
	} else if in.TransactionType == signproxy.TransactionType_PreSign {
		if entResp.Status != transaction.StatusPendingSigninfo {
			flagErr = 1
		}
		err = entResp.Update().
			SetStatus(transaction.StatusPendingSign).
			SetMutex(false).
			Exec(ctx)
	} else if in.TransactionType == signproxy.TransactionType_Signature {
		if entResp.Status != transaction.StatusPendingSign {
			flagErr = 1
		}
		err = entResp.Update().
			SetStatus(transaction.StatusPendingBroadcast).
			SetMutex(false).
			Exec(ctx)
	} else if in.TransactionType == signproxy.TransactionType_Broadcast {
		if entResp.Status != transaction.StatusPendingBroadcast {
			flagErr = 1
		}
		err = entResp.Update().
			SetStatus(transaction.StatusPendingConfirm).
			SetTransactionIDChain(in.TransactionIdChain).
			SetMutex(false).
			Exec(ctx)
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

func GetTransaction(ctx context.Context, in *trading.GetTransactionRequest) (resp *ent.Transaction, err error) {
	resp, err = db.Client().Transaction.Query().Where(
		transaction.TransactionIDInsite(in.TransactionIDInsite),
	).Only(ctx)
	return
}
