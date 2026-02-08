// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"beaver-server/app/friend/api/internal/config"
	"beaver-server/app/friend/rpc/friend"
	"beaver-server/app/notification/rpc/notification"
	coregorm "beaver-server/app/pkg/gorm"
	"beaver-server/app/pkg/redis"
	"beaver-server/app/pkg/version"
	"beaver-server/app/usercenter/rpc/usercenter"

	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UsercenterRpc usercenter.Usercenter
	FriendRpc     friend.Friend
	NotifyRpc     notification.Notification
	Version       *version.VersionGenerator
	DB            *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDB := coregorm.InitGorm(c.DB.DataSource)
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
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpc)),
		FriendRpc:     friend.NewFriend(zrpc.MustNewClient(c.FriendRpc)),
		NotifyRpc:     notification.NewNotification(zrpc.MustNewClient(c.NotifyRpc)),
		Version:       version.NewVersionGenerator(redisClient, mysqlDB),
		DB:            mysqlDB,
	}
}
