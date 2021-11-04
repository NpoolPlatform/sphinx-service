package api

import (
	"context"

	"github.com/NpoolPlatform/sphinx-service/message/npool"

	"github.com/NpoolPlatform/sphinx-service/pkg/core"

	"google.golang.org/protobuf/types/known/emptypb"
)

// https://github.com/grpc/grpc-go/issues/3794
// require_unimplemented_servers=false

func (s *Server) GetCoinInfos(ctx context.Context, req *npool.GetCoinInfosRequest) (cilist *npool.CoinInfoList, err error) {
	cilist, err = core.GetCoinInfos(ctx, req)
	return
}

func (s *Server) GetCoinInfo(ctx context.Context, req *npool.GetCoinInfoRequest) (cilist *npool.CoinInfoRow, err error) {
	cilist, err = core.GetCoinInfo(ctx, req)
	return
}


// 没写完的放下面


// 余额查询
func (s *Server) GetBalance(ctx context.Context, in *npool.GetBalanceRequest) (ret *AccountBalance, err error) {
	return nil, nil
}

// 转账 / 提现
func (s *Server) ApplyTransaction(ctx context.Context, in *npool.ApplyTransactionRequest) (ret *emptypb.errEmpty, error) {
	return nil, nil
}

// 签名服务接入点
func (s *Server) PortalSign(ctx context.Context, in *npool.PortalSignInit) (ret *IdentityProof, err error) {
	return nil, nil
}

// 代理服务接入点
func (s *Server) PortalWallet(ctx context.Context, in *npool.PortalWalletInit) (ret *IdentityProof, err error) {
	return nil, nil
}

// 账户交易查询
func (s *Server) GetTxJSON(ctx context.Context, in *npool.GetTxJSONRequest) (ret *AccountTxJSON, err error) {
	return nil, nil
}

// 交易状态查询
func (s *Server) GetInsiteTxStatus(ctx context.Context, in *npool.GetInsiteTxStatusRequest) (ret *GetInsiteTxStatusResponse, err error) {
	return nil, nil
}




// 在写的放尾部

// 创建账户
func (s *Server) RegisterAccount(context.Context, *RegisterAccountRequest) (*AccountAddress, error) {
	return nil, nil
}