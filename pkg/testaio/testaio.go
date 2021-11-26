package testaio

import (
	"os"
	"strconv"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/message/npool/trading"
	"github.com/go-resty/resty/v2"
)

var (
	InitAlready bool
	Host        string
	RestyClient *resty.Client
	AccountInfo = trading.EntAccount{
		CoinName: "FIL",
		Address:  "testaddresshere",
	}
	AddressFrom   = "t15kaul3frzthweyc44njodf7qjch7nbrhatmrrui"
	AddressTo     = "t13cjqm4dlj26huz4y7bobaqh2m542stvunznmiqa"
	AmountFloat64 = 0.01
	InsiteTxType  = "payment"
	CoinInfo      = coininfo.CoinInfo{
		ID:        "8fbcbdc2-25ea-4ff0-b049-9d2f4c8ab646",
		Enum:      1,
		PreSale:   false,
		Name:      "FIL",
		Unit:      "FIL",
		LogoImage: "",
	}
	AccountUUID         = "6ba7b812-9dad-80b4-11d1-00c04fd430c8"
	TransactionIDInsite = "test-tx-6ba7b812-80b4-9dad-11d1"
)

func UnifyRestyQuery(path string, body interface{}) (resp *resty.Response) {
	resp, err := RestyClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(Host + path)
	LogWhenError(err)
	return
}

func LogWhenError(err error) {
	if err != nil {
		logger.Sugar().Warn(err)
	}
}

func RunByGithub() bool {
	var err error
	runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION"))
	return err == nil && runByGithubAction
}
