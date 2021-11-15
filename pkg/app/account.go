package app

import (
	"context"

	"github.com/NpoolPlatform/sphinx-service/message/npool"
	// "github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
)

// 创建账户
func RegisterAccount(ctx context.Context, in *npool.RegisterAccountRequest) (resp *npool.AccountAddress, err error) {
	// 1. generate prikey with in.Uuid(string) and rand perms
	// 2. get aws secret from apollo, connect s3, upload (coin_id,key,address)
	// 3. return address
	resp = &npool.AccountAddress{
		CoinId:  0,
		Address: "",
		Uuid:    "",
	}
	return
}

// 余额查询
func GetBalance(ctx context.Context, in *npool.GetBalanceRequest) (resp *npool.AccountBalance, err error) {
	resp = &npool.AccountBalance{}
	return
}

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
