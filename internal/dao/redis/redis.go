package redis

import (
	"context"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/config"
	"github.com/go-redis/redis/v8"
	"strings"
	"sync"
)

var (
	rdb  *redis.Client
	once sync.Once
)

func New(config *config.Config) *redis.Client {
	host := strings.Join([]string{config.Redis.Host, config.Redis.Port}, ":")
	once.Do(func() {
		if config.Redis.Enable != 1 {
			rdb = nil
			return
		}
		rdb = redis.NewClient(&redis.Options{
			Addr:     host,
			Password: config.Redis.Password,
			DB:       config.Redis.Database,
		})
		var err error
		for i := 0; i < 3; i++ {
			if _, err = rdb.Ping(context.TODO()).Result(); err != nil {
				break
			}
		}
	})
	return rdb
}
