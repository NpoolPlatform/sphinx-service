package core

import (
	"context"

	"github.com/NpoolPlatform/sphinx-service/message/npool"
	"github.com/NpoolPlatform/sphinx-service/pkg/db"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/coininfo"
)

var client *ent.Client

func init() {
	client = db.Client()
}

// 查询全部币种
func GetCoinInfos(ctx context.Context, req *npool.GetCoinInfosRequest) (cilist *npool.CoinInfoList, err error) {
	ent_resp := []*ent.CoinInfo{}
	ent_resp, err = client.CoinInfo.Query().All(ctx)
	var tmp_cir []*npool.CoinInfoRow
	for _, row := range ent_resp {
		tmp_cir = append(tmp_cir, &npool.CoinInfoRow{
				Id: row.ID,
				Name: row.Name,
				Unit: row.Unit,
				NeedSigninfo: row.NeedSigninfo,
		})
	}
	cilist = &npool.CoinInfoList{
		List: tmp_cir,
	}
	return cilist, err
}

// 查询单个币种
func GetCoinInfo(ctx context.Context, req *npool.GetCoinInfoRequest) ( coin_info *npool.CoinInfoRow, err error) {
	ent_resp := &ent.CoinInfo{}
	ent_resp, err = client.CoinInfo.
	Query().
	Where(
			coininfo.ID(req.CoinId),
	).First(ctx)
	coin_info = &npool.CoinInfoRow{
		Id:           ent_resp.ID,
		NeedSigninfo: ent_resp.NeedSigninfo,
		Name:         ent_resp.Name,
		Unit:         ent_resp.Unit,
	}
	return coin_info, err
}

