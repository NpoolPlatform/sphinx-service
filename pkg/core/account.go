package core

import (
	"context"

	"github.com/NpoolPlatform/sphinx-service/message/npool"
	// "github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
	// "github.com/NpoolPlatform/sphinx-service/pkg/db/ent/keystore"
)

// *创建账户
// under construction
func RegisterAccount(ctx context.Context, in *npool.RegisterAccountRequest) (obj *npool.AccountAddress, err error) {
	// generate prikey with in.Uuid(string) and rand perms
	// get aws secret from apollo, connect s3, upload (coin_id,key,address)
	// return address
	obj = &npool.AccountAddress{
		CoinId:  0,
		Address: "",
		Uuid:    "",
	}
	return
}

