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

var txType2Status map[signproxy.TransactionType]transaction.Status

func init() {
	// Transaction State Machine
	txType2Status = make(map[signproxy.TransactionType]transaction.Status, 4)
	txType2Status[signproxy.TransactionType_TransactionNew] = transaction.StatusPendingProcess
	txType2Status[signproxy.TransactionType_PreSign] = transaction.StatusPendingSigninfo
	txType2Status[signproxy.TransactionType_Signature] = transaction.StatusPendingBroadcast
	txType2Status[signproxy.TransactionType_Broadcast] = transaction.StatusPendingConfirm
}

func CreateTransaction(ctx context.Context, in *trading.CreateTransactionRequest, needManualReview bool, txType transaction.Type) (info *ent.Transaction, err error) {
	// get coin info
	coinInfo, err := db.Client().CoinInfo.Query().Where(coininfo.Name(in.Info.CoinName)).Only(ctx)
	if err != nil {
		err = xerrors.Errorf(in.Info.CoinName+" coin not found %v", err)
		return
	}
	// check if exists
	info, err = GetSameTransactionOrNil(ctx, in)
	if info != nil {
		return info, err
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
		Save(ctx)
	return info, err
}

func GetSameTransactionOrNil(ctx context.Context, in *trading.CreateTransactionRequest) (record *ent.Transaction, err error) {
	var info []*ent.Transaction
	info, err = db.Client().Transaction.Query().
		Where(
			transaction.And(
				transaction.TransactionIDInsite(in.Info.TransactionIDInsite),
			),
		).All(ctx)
	if err != nil {
		return
	}
	if len(info) > 0 {
		if len(info) > 1 {
			err = xerrors.New("impossible")
		}
		coinInfo, err := info[0].QueryCoin().Only(ctx)
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

func UpdateTransactionStatusDeprecated(ctx context.Context, in *trading.ACKRequest) (isSuccess bool, err error) {
	// half state machine
	txStatus, ok := txType2Status[in.TransactionType]
	if !ok {
		err = xerrors.Errorf("ack tx type incorrect, %v", in)
		return
	}
	err = db.Client().Transaction.Update().
		Where(
			transaction.TransactionIDInsite(in.TransactionIdInsite),
		).
		SetStatus(txStatus).
		SetTransactionIDChain(in.TransactionIdChain).
		SetMutex(false).
		SetUpdatetimeUtc(time.Now().UTC().Unix()).
		Exec(ctx)
	isSuccess = (err == nil)
	return
}

func UpdateTransactionStatusV0(ctx context.Context, in *trading.ACKRequest) (err error) {
	// Judge status
	var txStatus transaction.Status
	if !in.IsOkay || in.TransactionIdChain == "" {
		txStatus = transaction.StatusErrorExpected
		logger.Sugar().Infof("[proxy] failure reported from proxy, %v", in)
	} else {
		txStatus = transaction.StatusDone
	}
	// Update db
	err = db.Client().Transaction.Update().
		Where(
			transaction.TransactionIDInsite(in.TransactionIdInsite),
		).
		SetStatus(txStatus).
		SetTransactionIDChain(in.TransactionIdChain).
		SetMutex(false).
		SetUpdatetimeUtc(time.Now().UTC().Unix()).
		Exec(ctx)
	return
}

func GetTransaction(ctx context.Context, in *trading.GetTransactionRequest) (resp *ent.Transaction, err error) {
	resp, err = db.Client().Transaction.Query().Where(
		transaction.TransactionIDInsite(in.TransactionIDInsite),
	).Only(ctx)
	return
}
