package app

import (
	"context"
	"sync"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/signproxy"
	"github.com/NpoolPlatform/message/npool/trading" //nolint
	"github.com/NpoolPlatform/sphinx-service/pkg/crud"
	"golang.org/x/xerrors"
)

// next-version: use channel to improve performance under high qps

var (
	mapACK              map[string]*trading.ACKRequest
	ackMutex            sync.Mutex
	ackListenIntervalMs int
	ackListenTimeoutMs  int
)

func init() {
	mapACK = make(map[string]*trading.ACKRequest)
	if ackListenIntervalMs <= 0 {
		ackListenIntervalMs = 80
	}
	if ackListenTimeoutMs <= 0 {
		ackListenTimeoutMs = 6000
	}
}

func ACK(ctx context.Context, in *trading.ACKRequest) (resp *trading.ACKResponse, err error) {
	resp = &trading.ACKResponse{}
	if in.TransactionType == signproxy.TransactionType_TransactionNew ||
		in.TransactionType == signproxy.TransactionType_PreSign ||
		in.TransactionType == signproxy.TransactionType_Signature ||
		in.TransactionType == signproxy.TransactionType_Broadcast {
		err = crud.UpdateTransactionStatusV0(ctx, in)
		resp.IsOkay = (err == nil)
	} else {
		// remove these async logic in next-version
		ackMutex.Lock()
		mapACK[in.TransactionIdInsite] = in
		ackMutex.Unlock()
		resp.IsOkay = true
	}
	return
}

func ListenTillSucceeded(transactionIDInsite string) (val *trading.ACKRequest, err error) {
	logger.Sugar().Infof("[listener] listening on TID: %v", transactionIDInsite)
	var ok bool
	ackListenTimeoutMsLoop := ackListenTimeoutMs
	for !ok && ackListenTimeoutMsLoop > 0 {
		time.Sleep(time.Duration(ackListenIntervalMs) * time.Millisecond)
		ackListenTimeoutMsLoop -= ackListenIntervalMs
		ackMutex.Lock()
		val, ok = mapACK[transactionIDInsite]
		ackMutex.Unlock()
	}
	if val != nil {
		if !val.IsOkay {
			err = xerrors.New("tx rejected by proxy")
		}
		delete(mapACK, transactionIDInsite)
		logger.Sugar().Infof("[listener] %v got resp and return %v", transactionIDInsite, val)
	} else {
		err = xerrors.New("request timeout, please try again later")
	}
	return
}
