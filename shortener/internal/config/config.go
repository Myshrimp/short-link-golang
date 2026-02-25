package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	CacheRedis cache.CacheConf

	ShortUrlDB struct {
		DSN string //data source name
	}

	Sequence struct {
		DSN string //data source name
	}

	ShortUrlBlacklist []string
	Domain            string
}
