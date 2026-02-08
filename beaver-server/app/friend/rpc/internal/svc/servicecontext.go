package svc

import (
	"beaver-server/app/friend/model"
	"beaver-server/app/friend/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	FriendModel       model.FriendModel
	FriendverifyModel model.FriendverifyModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config:            c,
		FriendModel:       model.NewFriendModel(sqlConn, c.Cache),
		FriendverifyModel: model.NewFriendverifyModel(sqlConn, c.Cache),
	}
}
