package core

import (
	"github.com/NpoolPlatform/sphinx-service/pkg/db"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
)

var client *ent.Client

func init() {
	client = db.Client()
}

// 查询全部币种
func (s *Server) GetCoinInfos(ctx context.Context, req *GetCoinInfosRequest) (cilist *CoinInfoList, err error) {
	cilist, err = client.CoinInfo.
		Query().
		All(context)
	return cilist, err
}
