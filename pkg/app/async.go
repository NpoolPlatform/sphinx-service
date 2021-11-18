package app

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/price"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/sphinx-service/pkg/db"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	"github.com/NpoolPlatform/sphinx-service/pkg/message/message"
	"github.com/NpoolPlatform/sphinx-service/pkg/message/server"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

var ctxPublic context.Context

func init() {
	ctxPublic = context.Background()
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

// 审核服务 进行审核 TODO
func LetApproveTransaction(tx *ent.Transaction) (err error) {
	// approve result override
	_, err = db.Client().Transaction.Update().
		SetMutex(false).
		SetStatus(transaction.StatusPendingProcess).
		SetUpdatetimeUtc(time.Now().UTC().Unix()).
		Save(context.Background())
	if err != nil {
		err = status.Error(codes.Internal, "database error")
		return
	}
	go LetSendTransaction(tx)
	return
}

// 钱包代理 查询交易状态
func CheckIfTransactionComplete(tx *ent.Transaction) (err error) {
	return
}

// 钱包代理 进行转账
func LetSendTransaction(tx *ent.Transaction) {
	coinID, err := tx.QueryCoin().OnlyID(ctxPublic)
	if err != nil {
		logger.Sugar().Errorf("transaction no corresponding coin! %v", err)
		return
	}
	err = server.PublishDefaultNotification(&message.NotificationTransaction{
		CoinType:            sphinxplugin.CoinType(coinID),
		TransactionIDInsite: tx.TransactionIDInsite,
		AmountFloat64:       tx.AmountFloat64,
		AddressFrom:         tx.AddressFrom,
		AddressTo:           tx.AddressTo,
		TransactionIDChain:  tx.TransactionIDChain,
		SignatureUser:       tx.SignatureUser,
		SignaturePlatform:   tx.SignaturePlatform,
		CreatetimeUtc:       tx.CreatetimeUtc,
		UpdatetimeUtc:       tx.UpdatetimeUtc,
		IsSuccess:           false,
		IsFailed:            false,
	})
	if err != nil {
		logger.Sugar().Errorf("failed to send transaction to proxy: %v", err)
		return
	}
	_, err = tx.Update().
		SetMutex(true).
		SetUpdatetimeUtc(time.Now().UTC().Unix()).
		Save(ctxPublic)
	if err != nil {
		logger.Sugar().Errorf("[!WARNING!] db update failed: %v", err)
	}
}
