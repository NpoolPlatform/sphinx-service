package fil

import (
	"context"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/myxtype/filecoin-client"
)

var Client *filecoin.Client

func lotusURL(host string) string {
	return fmt.Sprintf("http://%v:1234/rpc/v0", host)
}

func SetHostWithToken(str, token string) {
	Client = filecoin.NewClient(lotusURL(str), token)
}

func GetSignInfo(addr string) (json string, err error) { return }

func GetBalance(addr string) (str string, err error) {
	addrStd, err := address.NewFromString(addr)
	if err != nil {
		return
	}
	bal, err := Client.WalletBalance(context.Background(), addrStd)
	str = bal.String()
	return
}

func BroadcastScript(transactionScript string) (transactionIdChain string, err error) {
	return
}
