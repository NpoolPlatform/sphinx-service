package crud

import (
	"context"

	"github.com/NpoolPlatform/sphinx-service/pkg/db"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/coininfo"
	"golang.org/x/xerrors"
)

var ctxPublic context.Context

func init() {
	ctxPublic = context.Background()
}

func CoinName2CoinID(ctx context.Context, name string) (id int32, err error) {
	entResp, err := db.Client().CoinInfo.Query().Where(coininfo.Name(name)).All(ctxPublic)
	if err != nil {
		err = xerrors.Errorf("[WARNING] cannot query coininfo, %v", err)
		return
	} else if len(entResp) == 0 {
		err = xerrors.New("didn't find any id with input coin: " + name)
	}
	id = entResp[0].CoinTypeID
	return
}

func CoinName2Coin(ctx context.Context, name string) (coinInfo *ent.CoinInfo, err error) {
	coinInfo, err = db.Client().CoinInfo.Query().Where(coininfo.Name(name)).Only(ctxPublic)
	return
}
