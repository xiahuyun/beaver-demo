// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UsercenterRpc zrpc.RpcClientConf
	FriendRpc     zrpc.RpcClientConf
	NotifyRpc     zrpc.RpcClientConf

	Cache struct {
		Addr     string
		Password string
		DB       int
	}

	DB struct {
		DataSource string
	}
}
