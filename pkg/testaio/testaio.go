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
