package crud

import (
	"context"
	"testing"
	"time"

	//nolint
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/signproxy"
	"github.com/NpoolPlatform/message/npool/trading"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	testinit "github.com/NpoolPlatform/sphinx-service/pkg/test-init"
	"github.com/NpoolPlatform/sphinx-service/pkg/testaio"
	"github.com/stretchr/testify/assert"
)

func TestCRUD(t *testing.T) {
	if testaio.RunByGithub() {
		return
	}
	if !testaio.InitAlready {
		assert.Nil(t, testinit.Init())
		testaio.InitAlready = true
	}
	ctx := context.Background()
	// test create
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
	testaio.AbortWhenError(t, err)
	// test check
	tx, err = GetSameTransactionOrNil(ctx, createTxRequest)
	testaio.AbortWhenError(t, err)
	logger.Sugar().Infof("GetSameTransactionOrNil result: %v", tx)
	// test update
	updateTxRequest := &trading.ACKRequest{
		TransactionType:     signproxy.TransactionType_Broadcast,
		CoinTypeId:          testaio.CoinInfo.Enum,
		TransactionIdInsite: testaio.TransactionIDInsite,
		TransactionIdChain:  "mocked-transaction-id-chain",
		Address:             testaio.AddressFrom, // no need
		Balance:             0,                   // no need
		IsOkay:              true,
		ErrorMessage:        "",
	}
	err = UpdateTransactionStatusV0(ctx, updateTxRequest)
	logger.Sugar().Infof("UpdateTransactionStatus result: %v", err)
	testaio.AbortWhenError(t, err)
	_, err = UpdateTransactionStatusDeprecated(ctx, updateTxRequest)
	logger.Sugar().Infof("UpdateTransactionStatusDeprecated result: %v", err)
	tx, err = GetTransaction(ctx, &trading.GetTransactionRequest{
		TransactionIDInsite: testaio.TransactionIDInsite,
	})
	logger.Sugar().Infof("GetTransaction result: %v", tx)
	testaio.AbortWhenError(t, err)
	assert.Nil(t, err)
}
