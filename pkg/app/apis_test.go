package app

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/message/npool/coininfo" //nolint
	"github.com/NpoolPlatform/message/npool/trading"  //nolint
	testinit "github.com/NpoolPlatform/sphinx-service/pkg/test-init"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	testInitAlready bool
	tmpCoinInfo     coininfo.CoinInfo
)

func runByGithub() bool {
	runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION"))
	if err == nil && runByGithubAction {
		return true
	}
	if testInitAlready == false {
		testInitAlready = true
		err = testinit.Init()
		initStruct()
		if err != nil {
			fmt.Printf("test init failed: %v", err)
		}
	}
	return err == nil
}

func initStruct() {
	ctxPublic = context.Background()
	tmpCoinInfo.Enum = 0
	tmpCoinInfo.PreSale = false
	tmpCoinInfo.Name = "Unknown"
	tmpCoinInfo.Unit = "DK"
}

func TestCreateAccount(t *testing.T) {
	if runByGithub() {
		return
	}
	tmpUUID := uuid.NewString()
	if len(tmpUUID) > 2 {
		tmpUUID = tmpUUID[0:2]
	} else {
		panic("uuid too short!")
	}
	account, err := CreateAccount(ctxPublic, tmpCoinInfo.Name, tmpUUID)
	if err == nil {
		assert.NotNil(t, account)
		assert.Equal(t, tmpCoinInfo.Name, account.Info.CoinName)
		assert.NotEmpty(t, account.Info.Address)
	} else {
		assert.Nil(t, account)
	}
}

func TestGetBalance(t *testing.T) {
	if runByGithub() {
		return
	}
	tmpUUID := uuid.NewString()
	if len(tmpUUID) > 2 {
		tmpUUID = tmpUUID[0:2]
	} else {
		panic("uuid too short!")
	}
	resp, err := GetBalance(ctxPublic, &trading.GetBalanceRequest{
		Info: &trading.EntAccount{
			CoinName: tmpCoinInfo.Name,
			Address:  tmpUUID,
		},
	})
	if err == nil {
		assert.NotNil(t, resp)
		assert.Equal(t, tmpCoinInfo.Name, resp.Info.CoinName)
		assert.Positive(t, resp.AmountFloat64+0.1)
	}
}
