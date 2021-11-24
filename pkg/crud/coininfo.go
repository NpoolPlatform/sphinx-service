package crud

import (
	"context"

	"github.com/NpoolPlatform/sphinx-service/pkg/db"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/coininfo"
)

var ctxPublic context.Context

func init() {
	ctxPublic = context.Background()
}

func CoinName2Coin(ctx context.Context, name string) (coinInfo *ent.CoinInfo, err error) {
	coinInfo, err = db.Client().CoinInfo.Query().Where(coininfo.Name(name)).Only(ctxPublic)
	return
}
