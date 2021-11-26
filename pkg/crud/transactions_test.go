package crud

import (
	"context"
	"testing"
	"time"

	//nolint
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/trading"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	"github.com/NpoolPlatform/sphinx-service/pkg/testaio"
	"github.com/stretchr/testify/assert"
)

func TestCRUD(t *testing.T) {
	ctx := context.Background()
	createTxRequest := &trading.CreateTransactionRequest{
		Info: &trading.BaseTx{
			TransactionIDInsite: testaio.TransactionIDInsite,
			CoinName:            testaio.CoinInfo.Name,
			AmountFloat64:       testaio.AmountFloat64,
			AddressFrom:         testaio.AddressFrom,
			AddressTo:           testaio.AddressTo,
			InsiteTxType:        testaio.InsiteTxType,
			CreatetimeUTC:       time.Now().UTC().Unix(),
		},
		UUIDSignature: "",
	}
	tx, err := CreateTransaction(ctx, createTxRequest, false, transaction.TypePayment)
	logger.Sugar().Infof("CreateTransaction result: %v", tx)
	testaio.LogWhenError(err)
	tx, err = GetSameTransactionOrNil(ctx, createTxRequest)
	logger.Sugar().Infof("GetSameTransactionOrNil result: %v", tx)
	testaio.LogWhenError(err)
	tx, err = GetTransaction(ctx, &trading.GetTransactionRequest{
		TransactionIDInsite: testaio.TransactionIDInsite,
	})
	logger.Sugar().Infof("GetTransaction result: %v", tx)
	testaio.LogWhenError(err)
	assert.Nil(t, err)
}
