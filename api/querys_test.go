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
)

var (
	tmpCoinInfo            coininfo.CoinInfo
	tmpAccountInfo         trading.CreateAccountResponse
	tmpTransactionIDInsite string
	testInitAlready        bool
	testHost               string
	RestyClient            *resty.Client
)

func TestWholeProcedure(t *testing.T) {
	if runByGithub() {
		return
	}
	// test create account
	// go MockAccountCreated()
	address := tCreateAccount()
	logger.Sugar().Infof("create account result: %v", address)
	// test get balance
	// go MockAccountBalance()
	if false {
		logger.Sugar().Infof("get balance result: %v", tGetBalance(address))
	}
	// test create transaction
	// go MockTransactionComplete()
	if false {
		logger.Sugar().Infof("create transaction result: %v", tCreateTransaction(address, address))
	}
}

func UnifyRestyQuery(path string, body interface{}) (resp *resty.Response) {
	resp, err := RestyClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(testHost + path)
	LogError(err)
	return
}

func tCreateAccount() string {
	body := &trading.CreateAccountRequest{
		CoinName: tmpCoinInfo.Name,
		UUID:     tmpAccountInfo.Info.UUID,
	}
	path := "/v1/create/wallet"
	resp := UnifyRestyQuery(path, body)
	expectedReturn := &trading.CreateAccountResponse{}
	err := json.Unmarshal(resp.Body(), expectedReturn)
	if err != nil {
		panic(resp.String())
	}
	tmpAccountInfo.Info.Address = expectedReturn.Info.Address
	return expectedReturn.Info.Address
}

func tCreateTransaction(addressFrom, addressTo string) (info string) {
	body := &trading.CreateTransactionRequest{
		CoinName:            tmpCoinInfo.Name,
		TransactionIDInsite: tmpTransactionIDInsite,
		AddressFrom:         addressFrom,
		AddressTo:           addressTo,
		AmountFloat64:       123456.789,
		InsiteTxType:        "payment",
		UUIDSignature:       "",
		CreatetimeUTC:       time.Now().UTC().Unix(),
	}
	path := "/v1/create/transaction"
	resp := UnifyRestyQuery(path, body)
	expectedReturn := &trading.CreateTransactionResponse{}
	err := json.Unmarshal(resp.Body(), expectedReturn)
	if err != nil {
		panic(resp.String())
	}
	return expectedReturn.Info
}

func tGetBalance(address string) (balance float64) {
	body := &trading.GetBalanceRequest{
		CoinName:     "Unknown",
		Address:      address,
		TimestampUTC: time.Now().UTC().Unix(),
	}
	path := "/v1/get/wallet/balance"
	resp := UnifyRestyQuery(path, body)
	expectedReturn := &trading.GetBalanceResponse{}
	err := json.Unmarshal(resp.Body(), expectedReturn)
	if err != nil {
		panic(resp.String())
	}
	return expectedReturn.Info.AmountFloat64
}

func tACK(req *trading.ACKRequest) (isOkay bool) {
	body := req
	path := "/v1/internal/ack"
	resp := UnifyRestyQuery(path, body)
	expectedReturn := trading.ACKResponse{}
	err := json.Unmarshal(resp.Body(), &expectedReturn)
	if err != nil {
		panic(resp.String())
	}
	return expectedReturn.IsOkay
}

func MockAccountCreated() (isOkay bool) {
	time.Sleep(300 * time.Millisecond)
	body := &trading.ACKRequest{
		TransactionType:     signproxy.TransactionType_WalletNew,
		CoinTypeId:          tmpCoinInfo.Enum,
		TransactionIdInsite: tmpAccountInfo.Info.UUID + tmpAccountInfo.Info.CoinName,
		TransactionIdChain:  "",
		Address:             "testaddresshere",
		Balance:             0.00,
		IsOkay:              true,
		ErrorMessage:        "",
	}
	isOkay = tACK(body)
	return
}

func MockAccountBalance() (isOkay bool) {
	time.Sleep(300 * time.Millisecond)
	body := &trading.ACKRequest{
		TransactionType:     signproxy.TransactionType_Balance,
		CoinTypeId:          tmpCoinInfo.Enum,
		TransactionIdInsite: "balance-" + tmpCoinInfo.Name + "-" + "testaddresshere",
		TransactionIdChain:  "",
		Address:             "testaddresshere",
		Balance:             0.00,
		IsOkay:              true,
		ErrorMessage:        "",
	}
	isOkay = tACK(body)
	return
}

func MockTransactionComplete() (isOkay bool) {
	time.Sleep(300 * time.Millisecond)
	body := &trading.ACKRequest{
		TransactionType:     signproxy.TransactionType_TransactionNew,
		CoinTypeId:          tmpCoinInfo.Enum,
		TransactionIdInsite: tmpTransactionIDInsite,
		TransactionIdChain:  "testchainidhere",
		Address:             "testaddresshere",
		Balance:             0.00,
		IsOkay:              true,
		ErrorMessage:        "",
	}
	isOkay = tACK(body)
	return
}

func init() {
	err := logger.Init(logger.InfoLevel, "/tmp/sphinx-service.log")
	LogError(err)
	if runByGithub() {
		return
	}
	tmpCoinInfo.Enum = 0
	tmpCoinInfo.ID = "6ba7b812-9dad-11d1-80b4-00c04fd430c8"
	tmpCoinInfo.PreSale = false
	tmpCoinInfo.Name = "Unknown"
	tmpCoinInfo.Unit = "DK"
	tmpAccountInfo = trading.CreateAccountResponse{
		Info: &trading.EntAccount{
			CoinName: "Unknown",
			UUID:     "6ba7b812-9dad-80b4-11d1-00c04fd430c8",
		},
	}
	tmpTransactionIDInsite = "test-tx-6ba7b812-80b4-9dad-11d1"
	testHost = "http://localhost:50160"
	RestyClient = resty.New()
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

func LogError(err error) {
	if err != nil {
		logger.Sugar().Warn(err)
	}
}
