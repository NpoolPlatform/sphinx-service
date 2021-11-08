package fil

import (
	"context"
	"encoding/json"
	"time"

	"github.com/NpoolPlatform/sphinx-service/message/agents"
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
