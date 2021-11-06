package fil

import (
	"context"
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/myxtype/filecoin-client"
)

var Client *filecoin.Client

func lotusRpcUrl(host string) string {
	return fmt.Sprintf("http://%v:1234/rpc/v0", host)
}

func SetHostWithToken(str string, token string) {
	Client = filecoin.NewClient(lotusRpcUrl(str), token)
}

func GetSignInfo(address string) (json string, err error) { return }

func GetBalance(addr string) (string, error) {
	addrStd, _ := address.NewFromString(addr)
	bal, err := Client.WalletBalance(context.Background(), addrStd)
	return bal.String(), err
}
