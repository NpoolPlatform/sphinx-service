package api

import (
	"context"

	"github.com/NpoolPlatform/sphinx-service/message/npool"

	"google.golang.org/protobuf/types/known/emptypb"
)

// 没写完的放下面

// 余额查询
func (s *Server) GetBalance(ctx context.Context, in *npool.GetBalanceRequest) (ret *npool.AccountBalance, err error) {
	return
}

// 转账 / 提现
func (s *Server) ApplyTransaction(ctx context.Context, in *npool.ApplyTransactionRequest) (ret *emptypb.Empty, err error) {
	return
}

// 账户交易查询
func (s *Server) GetTxJSON(ctx context.Context, in *npool.GetTxJSONRequest) (ret *npool.AccountTxJSON, err error) {
	return nil, nil
}

// 交易状态查询
func (s *Server) GetInsiteTxStatus(ctx context.Context, in *npool.GetInsiteTxStatusRequest) (ret *npool.GetInsiteTxStatusResponse, err error) {
	return nil, nil
}

// 在写的放尾部

// 创建账户
func (s *Server) RegisterAccount(context.Context, *npool.RegisterAccountRequest) (*npool.AccountAddress, error) {
	return nil, nil
}
