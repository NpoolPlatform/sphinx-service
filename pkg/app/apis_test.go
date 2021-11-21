package app

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/message/npool/coininfo"
	"github.com/NpoolPlatform/message/npool/trading"
	testinit "github.com/NpoolPlatform/sphinx-service/pkg/test-init"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	testInitAlready bool
	tmpCoinInfo     coininfo.CoinInfoRow
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
	tmpCoinInfo.CoinType = 0
	tmpCoinInfo.IsPresale = false
	tmpCoinInfo.Name = "Unknown"
	tmpCoinInfo.Unit = "DK"
}

func TestRegisterAccount(t *testing.T) {
	if runByGithub() {
		return
	}
	tmpUUID := uuid.NewString()
	if len(tmpUUID) > 2 {
		tmpUUID = tmpUUID[0:2]
	} else {
		panic("uuid too short!")
	}
	account, err := RegisterAccount(ctxPublic, tmpCoinInfo.Name, tmpUUID)
	if err == nil {
		assert.NotNil(t, account)
		assert.Equal(t, tmpCoinInfo.Name, account.CoinName)
		assert.NotEmpty(t, account.Address)
		assert.Equal(t, tmpUUID, account.Uuid)
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
		CoinName:     tmpCoinInfo.Name,
		Address:      tmpUUID,
		TimestampUtc: time.Now().UTC().Unix(),
	})
	if err == nil {
		assert.NotNil(t, resp)
		assert.Equal(t, tmpCoinInfo.Name, resp.CoinName)
		assert.Positive(t, resp.TimestampUtc)
		assert.Positive(t, resp.AmountFloat64+0.1)
	}
}
