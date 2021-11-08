package fil

import (
	"context"
	"fmt"

	"github.com/cyvadra/filecoin-client/types"
	"github.com/filecoin-project/go-address"
)

var (
	MessageMethod  uint64  = 0
	MessageVersion uint64  = 0
	GasLimit       int64   = 0
	GasFeeCap      int64   = 10000
	GasPremium     int64   = 10000
	MaxFeeFloat    float64 = 0.001
)

// todo
func GetTxStatus(CID string) (json string, err error) { return }
func GetTxJSON(addr string) (json string, err error)  { return }

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

func BroadcastScript(s *types.SignedMessage) (err error) {
	// 消息广播
	mid, err := Client.MpoolPush(context.Background(), s)
	if err != nil {
		fmt.Println("消息广播失败")
		fmt.Println(err)
	} else {
		fmt.Println("消息发送成功，message id:")
		fmt.Println(mid.String())
	}
	return
}
