package fil

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/cyvadra/filecoin-client/types"
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
)

var (
	MessageMethod  uint64
	MessageVersion uint64
	GasLimit       int64
	GasFeeCap      int64 = 10000
	GasPremium     int64 = 10000
	MaxFeeFloat          = 0.001
)

// todo
func GetTxJSON(addr string) (json string, err error) { return }

func GetSignInfo(addr string) (signInfo *SignInfoFIL, err error) {
	addrStd, err := address.NewFromString(addr)
	if err != nil {
		return
	}
	nonce, err := Client.MpoolGetNonce(context.Background(), addrStd)
	signInfo = &SignInfoFIL{
		Nonce:       nonce,
		Version:     MessageVersion,
		GasLimit:    GasLimit,
		GasFeeCap:   GasFeeCap,
		GasPremium:  GasPremium,
		Method:      MessageMethod,
		MaxFeeFloat: MaxFeeFloat,
	}
	return
}

func GetBalance(addr string) (str string, err error) {
	addrStd, err := address.NewFromString(addr)
	if err != nil {
		return
	}
	bal, err := Client.WalletBalance(context.Background(), addrStd)
	str = bal.String()
	return
}

func BroadcastScript(s *types.SignedMessage) (cID string, err error) {
	// 消息广播
	mid, err := Client.MpoolPush(context.Background(), s)
	if err != nil {
		logger.Sugar().Errorf("消息广播失败, %s", err)
	} else {
		cID = mid.String()
	}
	return
}

func GetTxStatus(cID string) (msg *types.Message, err error) {
	CTID, err := cid.Decode(cID)
	if err != nil {
		return
	}
	msg, err = Client.ChainGetMessage(context.Background(), CTID)
	return
}
