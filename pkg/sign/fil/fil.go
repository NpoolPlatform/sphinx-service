package fil

import (
	"encoding/base64"
	"errors"

	"github.com/NpoolPlatform/sphinx-sign/pkg/storage"
	"github.com/cyvadra/filecoin-client/local"
	"github.com/cyvadra/filecoin-client/types"
	"github.com/filecoin-project/go-address"
)

func SignScript(msg *types.Message) (s *types.SignedMessage, err error) {
	// 获取私钥（目前明文）
	pk := storage.Get(msg.From.String())
	if len(pk) == 0 {
		err = errors.New("PrivateKey not found")
		return
	}
	pkb, err := base64.StdEncoding.DecodeString(pk)
	if err != nil {
		err = errors.New("Invalid private key")
		return
	}
	// 离线签名
	s, err = local.WalletSignMessage(types.KTBLS, pkb, msg)
	if err != nil {
		return
	}
	err = local.WalletVerifyMessage(s)
	return
}

func CreateAccount() (ki *types.KeyInfo, addr *address.Address, err error) {
	ki, addr, err = local.WalletNew(types.KTBLS)
	if err != nil {
		return
	}
	err = storage.Set(addr.String(), base64.StdEncoding.EncodeToString(ki.PrivateKey))
	return
}
