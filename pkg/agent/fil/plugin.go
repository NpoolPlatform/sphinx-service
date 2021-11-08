package fil

import (
	"context"
	"fmt"

	"github.com/cyvadra/filecoin-client/types"
	"github.com/filecoin-project/go-address"
)

// todo
func GetSignInfo(addr string) (json string, err error) { return }
func GetTxStatus(CID string) (json string, err error)  { return }
func GetTxJSON(addr string) (json string, err error)   { return }

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
