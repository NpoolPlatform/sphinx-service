package fil

import (
	"context"

	"github.com/NpoolPlatform/sphinx-service/message/agents"
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

// to be implemented
func (Server) SignScript(ctx context.Context, in *agents.SignScriptRequest) (resp *agents.SignScriptResponse, err error) {
	return
}
