package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NotificationEventModel = (*customNotificationEventModel)(nil)

type (
	// NotificationEventModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNotificationEventModel.
	NotificationEventModel interface {
		notificationEventModel
	}

	customNotificationEventModel struct {
		*defaultNotificationEventModel
	}
)

// NewNotificationEventModel returns a model for the database table.
func NewNotificationEventModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) NotificationEventModel {
	return &customNotificationEventModel{
		defaultNotificationEventModel: newNotificationEventModel(conn, c, opts...),
	}
}
