package fil

import (
	"fmt"
	"math/big"

	"github.com/cyvadra/filecoin-client"
)

var Client *filecoin.Client

type SignInfoFIL struct {
	Nonce       uint64
	Version     uint64
	GasLimit    int64
	GasFeeCap   int64
	GasPremium  int64
	Method      uint64
	MaxFeeFloat float64
}

func SetHostWithToken(str, token string) {
	Client = filecoin.NewClient(fmt.Sprintf("http://%v:1234/rpc/v0", str), token)
}

func DecomposeStringInt(str string) (amountInt int64, amountDigits int32, amountString string) {
	// get original value
	amountDigits = 18 // for initial result: x = x*10^18
	bi, _ := new(big.Int).SetString(str, 10)
	// divide by 10^9
	filExp := new(big.Int)
	filExp.Exp(big.NewInt(10), big.NewInt(9), nil)
	// make a float copy
	bf := new(big.Float).SetInt(bi)
	bf.Quo(bf, new(big.Float).SetInt64(1000000000))
	amountString = fmt.Sprintf("%f", bf)
	amountDigits = 9
	amountInt, _ = bf.Int64()
	return
}
