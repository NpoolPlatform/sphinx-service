package app

import (
	"time"

	"github.com/NpoolPlatform/message/npool/trading" //nolint
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MARK: optimization can be made, use channel (maybe) to improve performance under high qps payload

var mapACK map[string]*trading.ACKRequest

func init() {
	mapACK = make(map[string]*trading.ACKRequest)
}

func ACK(in *trading.ACKRequest) (resp *trading.ACKResponse) {
	mapACK[in.TransactionIdInsite] = in
	resp = &trading.ACKResponse{
		IsOkay: true,
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
			err = status.Error(codes.Internal, val.ErrorMessage)
		}
		mapACK[transactionIDInsite] = nil
	} else {
		err = status.Error(codes.Unavailable, "request timeout, please try again later")
	}
	return
}
