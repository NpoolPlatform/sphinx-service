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
func (s *Server) GetBalance(ctx context.Context, in *npool.GetBalanceRequest) (resp *npool.AccountBalance, err error) {
	resp, err = app.GetBalance(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("getbalance error: %f", err)
		resp = &npool.AccountBalance{}
		if DebugFlag {
			err = errInternal
		}
	}
	return
}

// 转账 / 提现
func (s *Server) ApplyTransaction(ctx context.Context, in *npool.ApplyTransactionRequest) (resp *npool.SuccessInfo, err error) {
	resp, err = app.ApplyTransaction(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("applytransaction error: %f", err)
		resp = &npool.SuccessInfo{}
		if DebugFlag {
			err = errInternal
		}
	}
	return
}

// TODO: 账户交易查询
func (s *Server) GetTxJSON(ctx context.Context, in *npool.GetTxJSONRequest) (resp *npool.AccountTxJSON, err error) {
	resp, err = app.GetTxJSON(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("gettxjson error: %f", err)
		resp = &npool.AccountTxJSON{}
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
func (s *Server) RegisterAccount(ctx context.Context, in *npool.RegisterAccountRequest) (resp *npool.AccountAddress, err error) {
	resp, err = app.RegisterAccount(ctx, in.CoinName, in.Uuid)
	if err != nil {
		logger.Sugar().Errorf("registeraccount error: %v", err)
		resp = &npool.AccountAddress{}
		if DebugFlag {
			err = errInternal
		}
	}
	return
}

// 接收异步返回
func (s *Server) ACK(_ context.Context, in *npool.ACKRequest) (*npool.ACKResponse, error) {
	return app.ACK(in), nil
}
