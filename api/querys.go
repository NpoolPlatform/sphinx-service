package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/trading" //nolint
	"github.com/NpoolPlatform/sphinx-service/pkg/app"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

var errInternal = status.Error(codes.Internal, "internal server error")

// 余额查询
func (s *Server) GetWalletBalance(ctx context.Context, in *npool.GetWalletBalanceRequest) (resp *npool.GetWalletBalanceResponse, err error) {
	resp, err = app.GetWalletBalance(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("GetWalletBalance error: %v", err)
		resp = &npool.GetWalletBalanceResponse{}
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
func (s *Server) GetTransaction(ctx context.Context, in *npool.GetTransactionRequest) (resp *npool.GetTransactionResponse, err error) {
	resp, err = app.GetTransaction(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("GetTransaction error: %v", err)
		resp = &npool.GetTransactionResponse{}
		if DebugFlag {
			err = errInternal
		}
	}
	return
}

// 创建账户
func (s *Server) CreateWallet(ctx context.Context, in *npool.CreateWalletRequest) (resp *npool.CreateWalletResponse, err error) {
	resp, err = app.CreateWallet(ctx, in.CoinName, in.UUID)
	if err != nil {
		logger.Sugar().Errorf("create account error: %v", err)
		resp = &npool.CreateWalletResponse{}
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

// ping pong
func (s *Server) Version(ctx context.Context, in *emptypb.Empty) (*npool.VersionResponse, error) {
	return &npool.VersionResponse{Info: "mvp"}, nil
}
