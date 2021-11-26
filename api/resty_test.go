package api

import (
	"encoding/json"
	"testing"
	"time"

	//nolint
	"github.com/NpoolPlatform/go-service-framework/pkg/logger" //nolint
	"github.com/NpoolPlatform/message/npool/signproxy"
	"github.com/NpoolPlatform/message/npool/trading"
	"github.com/NpoolPlatform/sphinx-service/pkg/testaio"
	resty "github.com/go-resty/resty/v2"
	"golang.org/x/xerrors"
)

func TestWholeProcedure(t *testing.T) {
	if testaio.RunByGithub() {
		return
	}
	// test create account
	address := tCreateWallet()
	logger.Sugar().Infof("create account result: %v", address)
	if address == "" {
		return
	}
	// test get balance
	logger.Sugar().Infof("get balance result: %v", tGetWalletBalance(address))
	// test create transaction
	logger.Sugar().Infof("create transaction result: %v", tCreateTransaction(address, address))
}

func tCreateWallet() string {
	body := &trading.CreateWalletRequest{
		CoinName: testaio.CoinInfo.Name,
		UUID:     testaio.AccountUUID,
	}
	path := "/v1/create/wallet"
	resp := testaio.UnifyRestyQuery(path, body)
	expectedReturn := &trading.CreateWalletResponse{}
	err := json.Unmarshal(resp.Body(), expectedReturn)
	if resp.StatusCode() != 200 {
		err := xerrors.New(resp.String())
		logger.Sugar().Error(err)
		return ""
	}
	if err != nil || expectedReturn.Info == nil {
		logger.Sugar().Errorf("unexpected response from proxy, %v %v", err, resp.String())
		return ""
	}
	testaio.AccountInfo.Address = expectedReturn.Info.Address
	return expectedReturn.Info.Address
}

func tCreateTransaction(addressFrom, addressTo string) (info string) {
	body := &trading.CreateTransactionRequest{
		Info: &trading.BaseTx{
			CoinName:            testaio.CoinInfo.Name,
			TransactionIDInsite: testaio.TransactionIDInsite,
			AddressFrom:         addressFrom,
			AddressTo:           addressTo,
			AmountFloat64:       123456.789,
			InsiteTxType:        "payment",
			CreatetimeUTC:       time.Now().UTC().Unix(),
		},
		UUIDSignature: "",
	}
	path := "/v1/create/transaction"
	resp := testaio.UnifyRestyQuery(path, body)
	if resp.StatusCode() != 200 {
		err := xerrors.New(resp.String())
		logger.Sugar().Error(err)
		return
	}
	expectedReturn := &trading.CreateTransactionResponse{}
	if resp.StatusCode() != 200 {
		err := xerrors.New(resp.String())
		logger.Sugar().Error(err)
		return ""
	}
	err := json.Unmarshal(resp.Body(), expectedReturn)
	if err != nil {
		logger.Sugar().Errorf("unexpected response from proxy, %v %v", err, resp.String())
	}
	return expectedReturn.Info.TransactionIDInsite
}

func tGetWalletBalance(address string) (balance float64) {
	body := &trading.GetWalletBalanceRequest{
		Info: &trading.EntAccount{
			CoinName: "Unknown",
			Address:  address,
		},
	}
	path := "/v1/get/wallet/balance"
	resp := testaio.UnifyRestyQuery(path, body)
	expectedReturn := &trading.GetWalletBalanceResponse{}
	err := json.Unmarshal(resp.Body(), expectedReturn)
	if resp.StatusCode() != 200 {
		err := xerrors.New(resp.String())
		logger.Sugar().Error(err)
		return
	}
	if err != nil || expectedReturn.Info == nil {
		logger.Sugar().Errorf("unexpected response from proxy, %v %v", err, resp.String())
	}
	return expectedReturn.AmountFloat64
}

func tACK(req *trading.ACKRequest) (isOkay bool) {
	body := req
	path := "/v1/internal/ack"
	resp := testaio.UnifyRestyQuery(path, body)
	expectedReturn := trading.ACKResponse{}
	err := json.Unmarshal(resp.Body(), &expectedReturn)
	if resp.StatusCode() != 200 {
		err := xerrors.New(resp.String())
		logger.Sugar().Error(err)
		return
	}
	if err != nil {
		logger.Sugar().Errorf("unexpected response from proxy, %v %v", err, resp.String())
	}
	return expectedReturn.IsOkay
}

func MockAccountCreated() (isOkay bool) {
	time.Sleep(300 * time.Millisecond)
	body := &trading.ACKRequest{
		TransactionType:     signproxy.TransactionType_WalletNew,
		CoinTypeId:          testaio.CoinInfo.Enum,
		TransactionIdInsite: testaio.AccountUUID + testaio.AccountInfo.CoinName,
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
		CoinTypeId:          testaio.CoinInfo.Enum,
		TransactionIdInsite: "balance-" + testaio.CoinInfo.Name + "-" + "testaddresshere",
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
		CoinTypeId:          testaio.CoinInfo.Enum,
		TransactionIdInsite: testaio.TransactionIDInsite,
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
	testaio.LogWhenError(err)
	if testaio.RunByGithub() {
		return
	}
	testaio.RestyClient = resty.New()
}
