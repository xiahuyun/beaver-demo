// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	WebSocket struct {
		DefaultReadLimit    int64
		DefaultReadTimeout  int64
		DefaultWriteTimeout int64
	}
}
