package fil

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/NpoolPlatform/sphinx-sign/message/agents"
	"github.com/cyvadra/filecoin-client/types"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
)

type Server agents.UnimplementedSignServer

func (Server) CreateAccount(ctx context.Context, in *agents.CreateAccountRequest) (resp *agents.CreateAccountResponse, err error) {
	_, addr, err := CreateAccount()
	if err != nil {
		return
	}
	resp = &agents.CreateAccountResponse{
		CoinId:  in.CoinId,
		Address: addr.String(),
		Uuid:    in.Uuid,
	}
	return
}

func (Server) SignScript(ctx context.Context, in *agents.SignScriptRequest) (resp *agents.SignScriptResponse, err error) {
	addrFrom, err := address.NewFromString(in.AddressFrom)
	if err != nil {
		return
	}
	addrTo, err := address.NewFromString(in.AddressTo)
	if err != nil {
		return
	}
	signInfoFIL := &SignInfoFIL{}
	err = json.Unmarshal([]byte(in.SignInfoJson), signInfoFIL)
	if err != nil {
		err = errors.New("SignInfoJSON format error, cannot Unmarshal")
		return
	}
	amount := StringValue2BigInt(in.AmountString)
	msg := &types.Message{
		Version:    signInfoFIL.Version,
		To:         addrTo,
		From:       addrFrom,
		Nonce:      signInfoFIL.Nonce,
		Value:      amount,
		GasLimit:   signInfoFIL.GasLimit,
		GasFeeCap:  abi.NewTokenAmount(signInfoFIL.GasFeeCap),
		GasPremium: abi.NewTokenAmount(signInfoFIL.GasPremium),
		Method:     0,
		Params:     []byte{},
	}
	signedMessage, err := SignScript(msg)
	if err != nil {
		return
	}
	signedJSON, err := json.Marshal(signedMessage)
	if err != nil {
		err = errors.New("Internal Server Error: invalid signedMessage json")
		return
	}
	resp = &agents.SignScriptResponse{
		CoinId:              in.CoinId,
		TransactionIdInsite: in.TransactionIdInsite,
		AddressFrom:         in.AddressFrom,
		AddressTo:           in.AddressTo,
		AmountInt:           in.AmountInt,
		AmountDigits:        in.AmountDigits,
		AmountString:        amount.String(),
		ScriptJson:          string(signedJSON),
	}
	return
}
