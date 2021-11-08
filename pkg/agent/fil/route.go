package fil

import (
	"context"
	"encoding/json"
	"time"

	"github.com/NpoolPlatform/sphinx-service/message/agents"
	"github.com/cyvadra/filecoin-client/types"
)

type Server agents.UnimplementedPluginServer

func (Server) GetSignInfo(ctx context.Context, in *agents.GetSignInfoRequest) (sio *agents.SignInfo, err error) {
	sioFIL, err := GetSignInfo(in.Address)
	if err != nil {
		return
	}
	js, err := json.Marshal(sioFIL)
	sio = &agents.SignInfo{
		Json: string(js),
	}
	return
}

func (Server) GetBalance(ctx context.Context, in *agents.GetBalanceRequest) (acb *agents.AccountBalance, err error) {
	acbStr, err := GetBalance(in.Address)
	if err != nil {
		return
	}
	amountInt, amountDigits, amountString := DecomposeStringInt(acbStr)
	acb = &agents.AccountBalance{
		CoinId:       0,
		Address:      in.Address,
		TimestampUtc: time.Now().UnixNano(),
		AmountInt:    amountInt,
		AmountDigits: amountDigits,
		AmountString: amountString,
	}
	return
}

func (Server) BroadcastScript(ctx context.Context, in *agents.BroadcastScriptRequest) (resp *agents.BroadcastScriptResponse, err error) {
	msg := &types.SignedMessage{}
	err = json.Unmarshal([]byte(in.TransactionScript), msg)
	if err != nil {
		return
	}
	cid, err := BroadcastScript(msg)
	resp = &agents.BroadcastScriptResponse{
		TransactionIdChain: cid,
	}
	return
}
