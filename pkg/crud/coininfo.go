package crud

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/sphinx-service/pkg/db"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/coininfo"
)

var ctxPublic context.Context

func init() {
	ctxPublic = context.Background()
}

func CoinName2CoinID(name string) (id int32, err error) {
	entResp, err := db.Client().CoinInfo.Query().Where(coininfo.Name(name)).Only(ctxPublic)
	if err != nil || entResp == nil {
		logger.Sugar().Warnf("didn't get coin id, err:", err)
		return
	}
	id = entResp.CoinTypeID
	return
}
