package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	
	ShortUrlDB struct {
		DSN string //data source name
	}

	Sequence struct {
		DSN string //data source name
	}

	ShortUrlBlacklist []string
	Domain string
}
