package app

import (
	"context"
	"testing"

	"github.com/NpoolPlatform/message/npool/trading" //nolint
	"github.com/NpoolPlatform/sphinx-service/pkg/testaio"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateWallet(t *testing.T) {
	if testaio.RunByGithub() {
		return
	}
	tmpUUID := uuid.NewString()
	if len(tmpUUID) > 2 {
		tmpUUID = tmpUUID[0:2]
	} else {
		panic("uuid too short!")
	}
	account, err := CreateWallet(context.Background(), testaio.CoinInfo.Name, tmpUUID)
	if err == nil {
		assert.NotNil(t, account)
		assert.Equal(t, testaio.CoinInfo.Name, account.Info.CoinName)
		assert.NotEmpty(t, account.Info.Address)
	} else {
		assert.Nil(t, account)
	}
}

func TestGetWalletBalance(t *testing.T) {
	if testaio.RunByGithub() {
		return
	}
	tmpUUID := uuid.NewString()
	if len(tmpUUID) > 2 {
		tmpUUID = tmpUUID[0:2]
	} else {
		panic("uuid too short!")
	}
	resp, err := GetWalletBalance(context.Background(), &trading.GetWalletBalanceRequest{
		Info: &trading.EntAccount{
			CoinName: testaio.CoinInfo.Name,
			Address:  tmpUUID,
		},
	})
	if err == nil {
		assert.NotNil(t, resp)
		assert.Equal(t, testaio.CoinInfo.Name, resp.Info.CoinName)
		assert.Positive(t, resp.AmountFloat64+0.1)
	}
}
