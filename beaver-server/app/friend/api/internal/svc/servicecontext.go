// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"beaver-server/app/friend/api/internal/config"
	"beaver-server/app/friend/rpc/friend"
	"beaver-server/app/usercenter/rpc/usercenter"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UsercenterRpc usercenter.Usercenter
	FriendRpc     friend.Friend
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpc)),
		FriendRpc:     friend.NewFriend(zrpc.MustNewClient(c.FriendRpc)),
	}
}
