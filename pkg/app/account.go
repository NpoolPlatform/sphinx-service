package app

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/trading"
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
