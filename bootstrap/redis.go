package bootstrap

import (
	"chinadaily_com_cn/pkg/config"
	"fmt"
	"github.com/gocolly/redisstorage"
	"sync"
)

var (
	Storage   *redisstorage.Storage
	onceRedis sync.Once
)

// SetupRedisStorage 初始化 Redis Storage
func SetupRedisStorage() {
	onceRedis.Do(func() {
		Storage = &redisstorage.Storage{
			Address: fmt.Sprintf("%s:%s",
				config.GetString("redis.storage.host"),
				config.GetString("redis.storage.port"),
			),
			Password: config.GetString("redis.storage.password"),
			DB:       config.GetInt("redis.storage.db"),
			Prefix:   config.GetString("redis.storage.prefix"),
		}
	})
}
