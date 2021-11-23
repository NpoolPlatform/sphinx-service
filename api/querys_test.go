package api

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"
	"time"

	//nolint
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/coininfo" //nolint
	"github.com/NpoolPlatform/message/npool/signproxy"
	"github.com/NpoolPlatform/message/npool/trading"
	testinit "github.com/NpoolPlatform/sphinx-service/pkg/test-init"
	resty "github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

var (
	tmpCoinInfo            coininfo.CoinInfo
	tmpAccountInfo         trading.CreateAccountResponse
	tmpTransactionIDInsite string
	testInitAlready        bool
	testHost               string
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	tmpCoinInfo.Enum = 0
	tmpCoinInfo.ID = "6ba7b812-9dad-11d1-80b4-00c04fd430c8"
	tmpCoinInfo.PreSale = false
	tmpCoinInfo.Name = "Unknown"
	tmpCoinInfo.Unit = "DK"
	tmpAccountInfo.CoinName = "Unknown"
	tmpAccountInfo.Uuid = "6ba7b812-9dad-80b4-11d1-00c04fd430c8"
	tmpTransactionIDInsite = "test-tx-6ba7b812-80b4-9dad-11d1"
	testHost = "http://localhost:50160"
}

func TestWholeProcedure(t *testing.T) {
	if runByGithub() {
		return
	}
	var err error
	// test create account
	go MockAccountCreated()
	err = tCreateAccount()
	assert.Nil(t, err)
	assert.NotEmpty(t, tmpAccountInfo.Address)
	// test get balance
	go MockAccountBalance()
	resp, err := tGetBalance(tmpAccountInfo.Address)
	LogError(err)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Zero(t, resp.AmountFloat64)
	// test create transaction
	// transaction would fail, but err should be nil
	go MockTransactionComplete()
	err = tCreateTransaction(tmpAccountInfo.Address, tmpAccountInfo.Address)
	LogError(err)
	assert.Nil(t, err)
}

func runByGithub() bool {
	var err error
	runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION"))
	if err == nil && runByGithubAction {
		return true
	}
	if testInitAlready == false {
		testInitAlready = true
		err = testinit.Init()
		LogError(err)
	}
	return false
}

func tCreateAccount() (err error) {
	cli := resty.New()
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(trading.CreateAccountRequest{
			CoinName: tmpCoinInfo.Name,
			Uuid:     tmpAccountInfo.Uuid,
		}).
		Post(testHost + "/v1/account/register")
	if err != nil {
		LogError(err)
		return
	}
	expectedReturn := &trading.CreateAccountResponse{}
	err = json.Unmarshal(resp.Body(), expectedReturn)
	if err != nil {
		LogError(err)
		return
	}
	tmpAccountInfo.Address = expectedReturn.Address
	return
}

func tCreateTransaction(addressFrom, addressTo string) (err error) {
	cli := resty.New()
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(trading.CreateTransactionRequest{
			CoinName:            tmpCoinInfo.Name,
			TransactionIdInsite: tmpTransactionIDInsite,
			AddressFrom:         addressFrom,
			AddressTo:           addressTo,
			AmountFloat64:       123456.789,
			Type:                "payment",
			UuidSignature:       "",
			CreatetimeUtc:       time.Now().UTC().Unix(),
		}).
		Post(testHost + "/v1/transaction/create")
	if err != nil {
		return
	}
	expectedReturn := &trading.CreateTransactionResponse{}
	err = json.Unmarshal(resp.Body(), expectedReturn)
	return
}

func tGetBalance(address string) (expectedReturn *trading.GetBalanceResponse, err error) {
	cli := resty.New()
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(trading.GetBalanceRequest{
			CoinName:     "Unknown",
			Address:      address,
			TimestampUtc: time.Now().UTC().Unix(),
		}).
		Post(testHost + "/v1/account/balance/get")
	if err != nil {
		return
	}
	expectedReturn = &trading.GetBalanceResponse{}
	err = json.Unmarshal(resp.Body(), expectedReturn)
	return
}

func tACK(req *trading.ACKRequest) (isOkay bool, err error) {
	isOkay = true
	cli := resty.New()
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(req).
		Post(testHost + "/v1/internal/ack")
	logger.Sugar().Warn(resp)
	LogError(err)
	expectedReturn := &trading.ACKResponse{}
	err = json.Unmarshal(resp.Body(), expectedReturn)
	isOkay = (err == nil)
	LogError(err)
	logger.Sugar().Warn("debug expectedReturn:")
	logger.Sugar().Warn(expectedReturn)
	return
}

func MockAccountCreated() (isOkay bool) {
	time.Sleep(200 * time.Millisecond)
	isOkay = true
	req := &trading.ACKRequest{
		TransactionType:     signproxy.TransactionType_WalletNew,
		CoinTypeId:          tmpCoinInfo.Enum,
		TransactionIdInsite: tmpAccountInfo.Uuid + tmpAccountInfo.CoinName,
		TransactionIdChain:  "",
		Address:             "testaddresshere",
		Balance:             0.00,
		IsOkay:              true,
		ErrorMessage:        "",
	}
	isOkay, err := tACK(req)
	LogError(err)
	return
}

func MockAccountBalance() (isOkay bool) {
	time.Sleep(200 * time.Millisecond)
	isOkay = true
	req := &trading.ACKRequest{
		TransactionType:     signproxy.TransactionType_Balance,
		CoinTypeId:          tmpCoinInfo.Enum,
		TransactionIdInsite: "balance-" + tmpCoinInfo.Name + "-" + "testaddresshere",
		TransactionIdChain:  "",
		Address:             "testaddresshere",
		Balance:             0.00,
		IsOkay:              true,
		ErrorMessage:        "",
	}
	isOkay, err := tACK(req)
	LogError(err)
	return
}

func MockTransactionComplete() (isOkay bool) {
	time.Sleep(200 * time.Millisecond)
	isOkay = true
	req := &trading.ACKRequest{
		TransactionType:     signproxy.TransactionType_TransactionNew,
		CoinTypeId:          tmpCoinInfo.Enum,
		TransactionIdInsite: tmpTransactionIDInsite,
		TransactionIdChain:  "testchainidhere",
		Address:             "testaddresshere",
		Balance:             0.00,
		IsOkay:              true,
		ErrorMessage:        "",
	}
	_, err := tACK(req)
	LogError(err)
	req.TransactionType = signproxy.TransactionType_PreSign
	isOkay, err = tACK(req)
	LogError(err)
	return
}

func LogError(err error) {
	if err != nil {
		logger.Sugar().Warn(err)
	}
}
