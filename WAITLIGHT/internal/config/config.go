// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	ShortUrlDB struct{
		DSN string
	}
	Sequence struct{
		DSN string
	}
	BaseString string
	// ShortUrlBlackList map[string]struct{}
	ShortUrlBlackList [] string
	ShortDomain string

	CacheRedis cache.CacheConf
}
