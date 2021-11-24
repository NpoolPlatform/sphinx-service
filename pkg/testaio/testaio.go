package testaio

import (
	"os"
	"strconv"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/message/npool/trading"
	"github.com/go-resty/resty/v2"
)

func init() {
	CoinInfo.Enum = 0
	CoinInfo.ID = "6ba7b812-9dad-11d1-80b4-00c04fd430c8"
	CoinInfo.PreSale = false
	CoinInfo.Name = "Unknown"
	CoinInfo.Unit = "DK"
	AccountInfo.CoinName = "Unknown"
	AccountUUID = "6ba7b812-9dad-80b4-11d1-00c04fd430c8"
	TransactionIDInsite = "test-tx-6ba7b812-80b4-9dad-11d1"
	Host = "http://localhost:50160"
}

var (
	CoinInfo            coininfo.CoinInfo
	AccountInfo         trading.EntAccount
	AccountUUID         string
	TransactionIDInsite string
	InitAlready         bool
	Host                string
	RestyClient         *resty.Client
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
