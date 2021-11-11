package fil

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/cyvadra/filecoin-client"
	"github.com/cyvadra/filecoin-client/local"
	"github.com/cyvadra/filecoin-client/types"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func init() {
	address.CurrentNetwork = address.Mainnet
	SetHostWithToken("172.16.30.117", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.ppK_nggwygh6kCPDlktdBtkGaqQXxoXM99iNx3-tZ9E")
}

func OfflineSign(ki *types.KeyInfo, msg *types.Message) (s *types.SignedMessage, err error) {
	// 离线签名
	s, err = local.WalletSignMessage(types.KTBLS, ki.PrivateKey, msg)
	if err != nil {
		fmt.Println("离线签名失败")
		fmt.Println(s)
		fmt.Println(err)
		return
	}
	fmt.Println("signed message: ")
	fmt.Println(s)
	fmt.Println(hex.EncodeToString(s.Signature.Data))
	// 验证签名
	if err = local.WalletVerifyMessage(s); err != nil {
		fmt.Println("验证签名失败", err)
	}
	return
}

func TestBroadcastScript(t *testing.T) {
	var err error
	// 静态设置
	toAddr, err := address.NewFromString("t1gvkap5jmv5k7gwpa42zj43i2oaai5zg74n66xra")
	pkStr := "c3pS5JcZEM1C5Yukor63mQ8DvADh1qQN/GrUsRA20XE="
	if err != nil {
		fmt.Println("收款地址错误", err)
		assert.Nil(t, err)
	}
	var fromAddr *address.Address
	var pk []byte
	pk, err = base64.StdEncoding.DecodeString(pkStr)
	if err != nil {
		fmt.Println("pk解码失败", err)
		assert.Nil(t, err)
	}
	// 设置key
	ki := &types.KeyInfo{
		Type:       types.KTBLS,
		PrivateKey: pk,
	}
	// 由key生成并确认地址
	fromAddr, err = local.WalletPrivateToAddress(crypto.SigTypeBLS, ki.PrivateKey)
	if err != nil {
		fmt.Println("生成地址失败", err)
		assert.Nil(t, err)
	}
	// 获取Nonce
	nonce, err := Client.MpoolGetNonce(context.Background(), *fromAddr)
	if err != nil {
		fmt.Println("获取Nonce失败，请检查主机配置")
		fmt.Println(err)
		return
	}
	msg := &types.Message{
		Version:    0,
		To:         toAddr,
		From:       *fromAddr,
		Nonce:      nonce,
		Value:      filecoin.FromFil(decimal.NewFromFloat(1.0001)),
		GasLimit:   0,
		GasFeeCap:  abi.NewTokenAmount(100),
		GasPremium: abi.NewTokenAmount(100),
		Method:     0,
		Params:     nil,
	}
	// 设置最大手续费
	maxFee := filecoin.FromFil(decimal.NewFromFloat(0.0001))
	// 估算GasLimit
	msg, err = Client.GasEstimateMessageGas(context.Background(), msg, &types.MessageSendSpec{MaxFee: maxFee}, nil)
	if err != nil {
		fmt.Println("GasEstimateMessageGas错误", err)
		assert.Nil(t, err)
	}
	signedMsg, err := OfflineSign(ki, msg)
	if err != nil {
		fmt.Println("签名失败", err)
		assert.Nil(t, err)
	}
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	s, err := BroadcastScript(signedMsg)
	assert.Nil(t, err)
	assert.NotEmpty(t, s)
}

func TestGetBalance(t *testing.T) {
	str, err := GetBalance("t1gvkap5jmv5k7gwpa42zj43i2oaai5zg74n66xra")
	assert.Nil(t, err)
	assert.NotEmpty(t, str)
	str, err = GetBalance("asdfasdfasdf")
	assert.NotNil(t, err)
	assert.Empty(t, str)
}

func TestGetTxStatus(t *testing.T) {
	msg, err := GetTxStatus("asdf")
	assert.NotNil(t, err)
	assert.Nil(t, msg)
	msg, err = GetTxStatus("bafy2bzacebcgpnxout6iu442fcybhli35wd4jh6uxdrqq4ajczgxgnzwe6vhi")
	assert.Nil(t, err)
	assert.NotNil(t, msg)
}

func TestGetSignInfo(t *testing.T) {
	signInfo, err := GetSignInfo("t3rhsc4yhuysbtl63kybatw7rmnii5lcn7w3rsitjnjbzezbvkk3cjnsabppog44u3acmv2rc3hfgvo7ubv6qq")
	assert.Nil(t, err)
	assert.NotNil(t, signInfo)
	signInfo, err = GetSignInfo("000000000000000000")
	assert.NotNil(t, err)
	assert.Nil(t, signInfo)
}

func TestGetTxJSON(t *testing.T) {
	json, err := GetTxJSON("to be implemented")
	assert.Nil(t, err)
	assert.Empty(t, json)
}
