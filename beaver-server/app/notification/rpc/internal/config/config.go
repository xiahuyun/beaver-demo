package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	CacheConf struct {
		Addr     string
		Password string
		DB       int
	}

	DB struct {
		DataSource string
	}

	Cache cache.CacheConf
}
