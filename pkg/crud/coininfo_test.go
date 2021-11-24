package crud

import (
	"context"
	"testing"

	//nolint
	"github.com/NpoolPlatform/sphinx-service/pkg/testaio"
	"github.com/stretchr/testify/assert"
)

func TestCoinName2CoinID(t *testing.T) {
	if testaio.RunByGithub() {
		return
	}
	id, err := CoinName2CoinID(context.Background(), "nil")
	assert.Zero(t, id)
	assert.NotNil(t, err)
}
