package svc

import (
	"shortener/internal/config"
	"shortener/model"
	"shortener/sequence"

	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	ShortUrlModel     model.ShortUrlMapModel
	Sequence          sequence.Sequence
	ShortUrlBlacklist map[string]struct{}
	Filter            *bloom.Filter
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.ShortUrlDB.DSN)
	blacklist := make(map[string]struct{})
	for _, word := range c.ShortUrlBlacklist {
		blacklist[word] = struct{}{}
	}

	store := redis.New(c.CacheRedis[0].Host, func(r *redis.Redis) {
		r.Type = redis.NodeType
	})

	filter := bloom.New(store, "bloom_filter", 20*1<<20)

	return &ServiceContext{
		Config:            c,
		ShortUrlModel:     model.NewShortUrlMapModel(conn, c.CacheRedis),
		Sequence:          sequence.NewMySQL(c.Sequence.DSN),
		ShortUrlBlacklist: blacklist,
		//Sequence: sequence.NewRedis(c.Sequence.RedisAddr, c.Sequence.RedisPassword, c.Sequence.RedisDB),
		Filter: filter,
	}
}
