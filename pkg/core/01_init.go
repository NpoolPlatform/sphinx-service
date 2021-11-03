package core

import (
	"github.com/NpoolPlatform/sphinx-service/message/npool"
	"github.com/NpoolPlatform/sphinx-service/pkg/db"
)

type Server struct {
	npool.UnimplementedTradingServer
}

var PendingServer *UnimplementedTradingServer

func init() {
	PendingServer := &UnimplementedTradingServer{}

}

func GetTradingServer() *TradingServer {
	return Server
}





// 没写完的放这里，写完删

// 创建账户
func (s *Server) RegisterAccount(context.Context, *RegisterAccountRequest) (*AccountAddress, error) {
	return nil, nil
}

// 余额查询
func (s *Server) GetBalance(context.Context, *GetBalanceRequest) (*AccountBalance, error) {
	return nil, nil
}

// 转账 / 提现
func (s *Server) ApplyTransaction(context.Context, *ApplyTransactionRequest) (*emptypb.Empty, error) {
	return nil, nil
}

// 签名服务接入点
func (s *Server) PortalSign(context.Context, *PortalSignInit) (*IdentityProof, error) {
	return nil, nil
}

// 代理服务接入点
func (s *Server) PortalWallet(context.Context, *PortalWalletInit) (*IdentityProof, error) {
	return nil, nil
}

// 账户交易查询
func (s *Server) GetTxJSON(context.Context, *GetTxJSONRequest) (*AccountTxJSON, error) {
	return nil, nil
}

// 交易状态查询
func (s *Server) GetInsiteTxStatus(context.Context, *GetInsiteTxStatusRequest) (*GetInsiteTxStatusResponse, error) {
	return nil, nil
}
