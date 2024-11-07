package Redis

import (
	"MyTest/Settings"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func Init(config *Settings.RedisConfig) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.Db,
	})
	Rdb = rdb
	if Rdb == nil {
		return fmt.Errorf("init redis fail")
	}
	fmt.Println("redis connect success")
	return nil
}
