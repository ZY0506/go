package redis

import (
	"LearningProject/web_app/settings"
	"fmt"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(redisconfig *settings.Redis) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			redisconfig.Host,
			redisconfig.Port,
		),
		Password: redisconfig.Password,
		DB:       redisconfig.DB,
		PoolSize: redisconfig.PoolSize,
	})
	return rdb.Ping().Err()
}

func Close() {
	_ = rdb.Close()
}
