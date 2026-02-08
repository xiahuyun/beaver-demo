package svc

import (
	"beaver-server/app/notification/model"
	"beaver-server/app/notification/rpc/internal/config"
	coregorm "beaver-server/app/pkg/gorm"
	"beaver-server/app/pkg/redis"
	"beaver-server/app/pkg/version"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config                 config.Config
	Version                *version.VersionGenerator
	NotificationEventModel model.NotificationEventModel
	NotificationInboxModel model.NotificationInboxModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDB := coregorm.InitGorm(c.DB.DataSource)
	redisClient, err := redis.NewClient(&redis.Options{
		Addr:     c.CacheConf.Addr,
		Password: c.CacheConf.Password,
		DB:       c.CacheConf.DB,
	})
	if err != nil {
		panic(err)
	}

	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config:                 c,
		Version:                version.NewVersionGenerator(redisClient, mysqlDB),
		NotificationEventModel: model.NewNotificationEventModel(sqlConn, c.Cache),
		NotificationInboxModel: model.NewNotificationInboxModel(sqlConn, c.Cache),
	}
}
