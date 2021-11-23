package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/trading" //nolint
	"github.com/NpoolPlatform/sphinx-service/pkg/app"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errInternal = status.Error(codes.Internal, "internal server error")

// 余额查询
func (s *Server) GetBalance(ctx context.Context, in *npool.GetBalanceRequest) (resp *npool.GetBalanceResponse, err error) {
	resp, err = app.GetBalance(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("getbalance error: %v", err)
		resp = &npool.GetBalanceResponse{}
		if DebugFlag {
			err = errInternal
		}
	}
	return
}

// 转账 / 提现
func (s *Server) CreateTransaction(ctx context.Context, in *npool.CreateTransactionRequest) (resp *npool.CreateTransactionResponse, err error) {
	resp, err = app.CreateTransaction(ctx, in)
	if err != nil {
		logger.Sugar().Error(err)
		resp = &npool.CreateTransactionResponse{}
		if DebugFlag {
			err = errInternal
		}
	}
	return
}

// 交易状态查询
func (s *Server) GetInsiteTxStatus(ctx context.Context, in *npool.GetInsiteTxStatusRequest) (resp *npool.GetInsiteTxStatusResponse, err error) {
	resp, err = app.GetInsiteTxStatus(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("getinsitetxstatus error: %v", err)
		resp = &npool.GetInsiteTxStatusResponse{}
		if DebugFlag {
			err = errInternal
		}
	}
	return
}

// 创建账户
func (s *Server) CreateAccount(ctx context.Context, in *npool.CreateAccountRequest) (resp *npool.CreateAccountResponse, err error) {
	resp, err = app.CreateAccount(ctx, in.CoinName, in.UUID)
	if err != nil {
		logger.Sugar().Errorf("create account error: %v", err)
		resp = &npool.CreateAccountResponse{}
		if DebugFlag {
			err = errInternal
		}
	}
	return
}

// 接收异步返回
func (s *Server) ACK(ctx context.Context, in *npool.ACKRequest) (*npool.ACKResponse, error) {
	return app.ACK(ctx, in)
}
