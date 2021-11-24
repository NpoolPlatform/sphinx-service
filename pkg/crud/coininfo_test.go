package crud

import (
	"context"
	"testing"

	//nolint
	"github.com/NpoolPlatform/sphinx-service/pkg/testaio"
	"github.com/stretchr/testify/assert"
)

func TestCoinName2Coin(t *testing.T) {
	if testaio.RunByGithub() {
		return
	}
	ctx := context.Background()
	coinInfo, err := CoinName2Coin(ctx, "nil")
	assert.Nil(t, coinInfo)
	assert.NotNil(t, err)
	coinInfo, err = CoinName2Coin(ctx, testaio.CoinInfo.Name)
	if err == nil {
		assert.NotNil(t, coinInfo)
		assert.Equal(t, coinInfo.Unit, testaio.CoinInfo.Unit)
		assert.Equal(t, coinInfo.CoinTypeID, testaio.CoinInfo.Enum)
	} else {
		assert.Nil(t, coinInfo)
	}
}
