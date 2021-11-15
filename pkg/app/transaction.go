package app

import (
	"context"

	"github.com/NpoolPlatform/sphinx-service/message/npool"
	// "github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
)

// 转账 / 提现
func ApplyTransaction(ctx context.Context, in *npool.ApplyTransactionRequest) (resp *npool.SuccessInfo, err error) {
	return
}

// TODO: 账户交易查询
func GetTxJSON(ctx context.Context, in *npool.GetTxJSONRequest) (resp *npool.AccountTxJSON, err error) {
	return
}

// 交易状态查询
func GetInsiteTxStatus(ctx context.Context, in *npool.GetInsiteTxStatusRequest) (resp *npool.GetInsiteTxStatusResponse, err error) {
	return
}

// 余额查询
func GetBalance(ctx context.Context, in *npool.GetBalanceRequest) (resp *npool.AccountBalance, err error) {
	resp = &npool.AccountBalance{}
	return
}
