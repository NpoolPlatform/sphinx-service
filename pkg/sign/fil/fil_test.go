package fil

import (
	"testing"

	"github.com/cyvadra/filecoin-client"
	"github.com/cyvadra/filecoin-client/local"
	"github.com/cyvadra/filecoin-client/types"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestSignScript(t *testing.T) {
	// (s *types.SignedMessage, err error)
	ki, addr, err := local.WalletNew(types.KTBLS)
	assert.Nil(t, err)
	msg := &types.Message{
		Version:    0,
		To:         *addr,
		From:       *addr,
		Nonce:      0,
		Value:      filecoin.FromFil(decimal.NewFromFloat(1.0001)),
		GasLimit:   0,
		GasFeeCap:  abi.NewTokenAmount(100),
		GasPremium: abi.NewTokenAmount(100),
		Method:     0,
		Params:     []byte{},
	}
	s, err := SignScript(ki, msg)
	assert.Nil(t, err)
	err = local.WalletVerifyMessage(s)
	assert.Nil(t, err)
}

func TestCreateAccount(t *testing.T) {
	// (ki *types.KeyInfo, addr *address.Address, err error)
	_, _, err := local.WalletNew(types.KTBLS)
	assert.Nil(t, err)
}
