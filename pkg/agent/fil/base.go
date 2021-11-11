package fil

import (
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/cyvadra/filecoin-client"
	"github.com/shopspring/decimal"
)

var Client *filecoin.Client

const priceScale = 1000000000000

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

func DecomposeStringUInt64(str string) (amountUint64 uint64) {
	// for initial result: x = x*10^18
	bi, _ := new(big.Int).SetString(str, 10)
	bf := new(big.Float).SetInt(bi)
	bf.Quo(bf, new(big.Float).SetInt64(1000000000000000000))
	bf64, _ := bf.Float64()
	amountUint64 = VisualPriceToDBPrice(bf64)
	return
}

func DecomposeStringFloat64(str string) (bf64 float64) {
	// for initial result: x = x*10^18
	bi, _ := new(big.Int).SetString(str, 10)
	bf := new(big.Float).SetInt(bi)
	bf.Quo(bf, new(big.Float).SetInt64(1000000000000000000))
	bf64, _ = bf.Float64()
	return
}

func VisualPriceToDBPrice(price float64) uint64 {
	myPrice := decimal.NewFromFloat(price).Mul(decimal.NewFromInt(priceScale))
	iPrice, err := strconv.ParseUint(myPrice.String(), 10, 64)
	if err != nil {
		return uint64(price * priceScale)
	}
	return iPrice
}

func DBPriceToVisualPrice(price uint64) float64 {
	myPrice := decimal.NewFromInt(int64(price)).Div(decimal.NewFromInt(priceScale))
	fPrice, err := strconv.ParseFloat(myPrice.String(), 64)
	if err != nil {
		return float64(price)
	}
	return fPrice
}

func RunByGithubAction() (isDebug bool) {
	runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION"))
	isDebug = err == nil && runByGithubAction
	return
}
