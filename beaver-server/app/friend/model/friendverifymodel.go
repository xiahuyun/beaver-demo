package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FriendverifyModel = (*customFriendverifyModel)(nil)

type (
	// FriendVerifyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFriendVerifyModel.
	FriendverifyModel interface {
		friendverifyModel
	}

	customFriendverifyModel struct {
		*defaultFriendverifyModel
	}
)

// NewFriendVerifyModel returns a model for the database table.
func NewFriendverifyModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) FriendverifyModel {
	return &customFriendverifyModel{
		defaultFriendverifyModel: newFriendverifyModel(conn, c, opts...),
	}
}
