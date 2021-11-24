package app

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/signproxy"
	"github.com/NpoolPlatform/message/npool/trading" //nolint
	"github.com/NpoolPlatform/sphinx-service/pkg/crud"
	"golang.org/x/xerrors"
)

// MARK: optimization can be made, use channel (maybe) to improve performance under high qps payload

var mapACK map[string]*trading.ACKRequest

func init() {
	mapACK = make(map[string]*trading.ACKRequest)
}

func ACK(ctx context.Context, in *trading.ACKRequest) (resp *trading.ACKResponse, err error) {
	resp = &trading.ACKResponse{
		IsOkay: false,
	}
	logger.Sugar().Warn(in)
	logger.Sugar().Warn(in.TransactionType)
	if in.TransactionType == signproxy.TransactionType_TransactionNew ||
		in.TransactionType == signproxy.TransactionType_PreSign ||
		in.TransactionType == signproxy.TransactionType_Signature ||
		in.TransactionType == signproxy.TransactionType_Broadcast {
		resp.IsOkay, err = crud.UpdateTransactionStatus(ctx, in)
	} else {
		mapACK[in.TransactionIdInsite] = in
		resp.IsOkay = true
	}
	return
}

func ListenTillSucceeded(transactionIDInsite string) (val *trading.ACKRequest, err error) {
	loopIntervalMs := 50
	timeoutMs := 6000
	var ok bool
	for !ok && timeoutMs > 0 {
		time.Sleep(time.Duration(loopIntervalMs) * time.Millisecond)
		timeoutMs -= loopIntervalMs
		val, ok = mapACK[transactionIDInsite]
	}
	if val != nil {
		if !val.IsOkay {
			err = xerrors.New("tx rejected by proxy")
		}
		val = &trading.ACKRequest{
			TransactionType:     val.TransactionType,
			CoinTypeId:          val.CoinTypeId,
			TransactionIdInsite: val.TransactionIdInsite,
			TransactionIdChain:  val.TransactionIdChain,
			Address:             val.Address,
			Balance:             val.Balance,
			IsOkay:              val.IsOkay,
			ErrorMessage:        val.ErrorMessage,
		}
		mapACK[transactionIDInsite] = nil
	} else {
		err = xerrors.New("request timeout, please try again later")
	}
	return
}
