package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NotificationInboxModel = (*customNotificationInboxModel)(nil)

type (
	// NotificationInboxModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNotificationInboxModel.
	NotificationInboxModel interface {
		notificationInboxModel
	}

	customNotificationInboxModel struct {
		*defaultNotificationInboxModel
	}
)

// NewNotificationInboxModel returns a model for the database table.
func NewNotificationInboxModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) NotificationInboxModel {
	return &customNotificationInboxModel{
		defaultNotificationInboxModel: newNotificationInboxModel(conn, c, opts...),
	}
}
