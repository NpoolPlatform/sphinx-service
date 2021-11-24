package app

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/signproxy"    //nolint
	"github.com/NpoolPlatform/message/npool/sphinxplugin" //nolint
	"github.com/NpoolPlatform/sphinx-service/pkg/db"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	"github.com/NpoolPlatform/sphinx-service/pkg/message/message"
	"github.com/NpoolPlatform/sphinx-service/pkg/message/server"
	"golang.org/x/xerrors"
)

// 进行审核
func MockApproveTransaction(tx *ent.Transaction) (err error) {
	// approve result override
	_, err = db.Client().Transaction.Update().
		SetMutex(false).
		SetStatus(transaction.StatusPendingProcess).
		SetUpdatetimeUtc(time.Now().UTC().Unix()).
		Save(context.Background())
	if err != nil {
		logger.Sugar().Warn(err)
		err = xerrors.Errorf("approval err %v", err)
		return
	}
	// TODO
	// send transaction when it be approved
	_, err = LetSendTransaction(tx)
	return
}

// 钱包代理 进行转账 done
func LetSendTransaction(tx *ent.Transaction) (txNew *ent.Transaction, err error) {
	// get coin info
	coinInfo, err := tx.QueryCoin().Only(context.Background())
	if err != nil {
		err = xerrors.Errorf("tx coin data invalid: %v", err)
		return
	}
	// send through rabbitmq to signproxy
	err = server.PublishDefaultNotification(&message.NotificationTransaction{
		CoinType:            sphinxplugin.CoinType(coinInfo.CoinTypeID),
		TransactionType:     signproxy.TransactionType_TransactionNew,
		TransactionIDInsite: tx.TransactionIDInsite,
		AmountFloat64:       tx.AmountFloat64,
		AddressFrom:         tx.AddressFrom,
		AddressTo:           tx.AddressTo,
		TransactionIDChain:  "", // no need
		SignatureUser:       tx.SignatureUser,
		SignaturePlatform:   tx.SignaturePlatform,
		CreatetimeUtc:       tx.CreatetimeUtc,
		UpdatetimeUtc:       tx.UpdatetimeUtc,
		IsSuccess:           false, // no need
		IsFailed:            false, // no need
	})
	if err != nil {
		err = xerrors.Errorf("failed to send transaction to proxy: %v", err)
		return
	}
	// update db status
	txNew, err = tx.Update().
		SetMutex(true).
		SetUpdatetimeUtc(time.Now().UTC().Unix()).
		Save(context.Background())
	if err != nil {
		err = xerrors.Errorf("db update failed: %v", err)
	}
	return txNew, err
}
