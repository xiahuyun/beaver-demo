// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"beaver-server/app/usercenter/api/internal/config"
	"beaver-server/app/usercenter/rpc/usercenter"

	"beaver-server/app/pkg/redis"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	Redis         *redis.Client
	UsercenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisClient, err := redis.NewClient(&redis.Options{
		Addr:     c.Cache.Addr,
		Password: c.Cache.Password,
		DB:       c.Cache.DB,
	})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:        c,
		Redis:         redisClient,
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpc)),
	}
}
