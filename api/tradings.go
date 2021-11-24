package api

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/trading" //nolint
	"github.com/NpoolPlatform/sphinx-service/pkg/app"
)

// 余额查询
func (s *Server) GetWalletBalance(ctx context.Context, in *npool.GetWalletBalanceRequest) (resp *npool.GetWalletBalanceResponse, err error) {
	resp, err = app.GetWalletBalance(ctx, in)
	if err != nil {
		err = PatchGRPCError(err, "get wallet balance")
		resp = &npool.GetWalletBalanceResponse{}
	}
	return
}

// 转账 / 提现
func (s *Server) CreateTransaction(ctx context.Context, in *npool.CreateTransactionRequest) (resp *npool.CreateTransactionResponse, err error) {
	resp, err = app.CreateTransaction(ctx, in)
	if err != nil {
		err = PatchGRPCError(err, "create transaction")
		resp = &npool.CreateTransactionResponse{}
	}
	return
}

// 交易状态查询
func (s *Server) GetTransaction(ctx context.Context, in *npool.GetTransactionRequest) (resp *npool.GetTransactionResponse, err error) {
	resp, err = app.GetTransaction(ctx, in)
	if err != nil {
		err = PatchGRPCError(err, "get transaction")
		resp = &npool.GetTransactionResponse{}
	}
	return
}

// 创建账户
func (s *Server) CreateWallet(ctx context.Context, in *npool.CreateWalletRequest) (resp *npool.CreateWalletResponse, err error) {
	resp, err = app.CreateWallet(ctx, in.CoinName, in.UUID)
	if err != nil {
		err = PatchGRPCError(err, "create wallet")
		resp = &npool.CreateWalletResponse{}
	}
	return
}
