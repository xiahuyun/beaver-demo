// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	Cache struct {
		Addr     string
		Password string
		DB       int
	}

	UsercenterRpc zrpc.RpcClientConf

	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
}
