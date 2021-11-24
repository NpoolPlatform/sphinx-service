package rules

import (
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent"
)

func GenerateTID4CreateWallet(coinInfo *ent.CoinInfo, uniqueByStr string) (transactionID string) {
	/*
		1、确定节点返回的唯一性
		2、不存
		3、可使返回结果相同方便测试
	*/
	transactionID = "001-" + coinInfo.ID.String()[0:9] + "-" + uniqueByStr
	return
}

func GenerateTID4GetWalletBalance(coinInfo *ent.CoinInfo, uniqueByStr string) (transactionID string) {
	/*
		1、确定节点返回的唯一性
		2、不存
		3、可使返回结果相同方便测试
	*/
	transactionID = "002-" + coinInfo.ID.String()[0:9] + "-" + uniqueByStr
	return
}
