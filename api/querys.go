package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/trading"
	"github.com/NpoolPlatform/sphinx-service/pkg/app"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var errInternal = status.Error(codes.Internal, "internal server error")

// 余额查询
func (s *Server) GetBalance(ctx context.Context, in *npool.GetBalanceRequest) (resp *npool.AccountBalance, err error) {
	resp, err = app.GetBalance(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("getbalance error: %w", err)
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
		logger.Sugar().Errorw("applytransaction error: %w", err)
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
		logger.Sugar().Errorw("gettxjson error: %w", err)
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
		logger.Sugar().Errorw("getinsitetxstatus error: %w", err)
		resp = &npool.GetInsiteTxStatusResponse{}
		if DebugFlag {
			err = errInternal
		}
	}
	return
}

// 创建账户
func (s *Server) RegisterAccount(ctx context.Context, in *npool.RegisterAccountRequest) (resp *npool.AccountAddress, err error) {
	resp, err = app.RegisterAccount(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("registeraccount error: %w", err)
		resp = &npool.AccountAddress{}
		if DebugFlag {
			err = errInternal
		}
	}
	return
}
