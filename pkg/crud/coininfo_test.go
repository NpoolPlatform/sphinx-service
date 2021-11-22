package crud

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/message/npool/coininfo"
	testinit "github.com/NpoolPlatform/sphinx-service/pkg/test-init"
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

func TestCoinName2CoinID(t *testing.T) {
	if runByGithub() {
		return
	}
	id, err := CoinName2CoinID("nil")
	assert.NotNil(t, err)
	assert.Zero(t, id)
}
