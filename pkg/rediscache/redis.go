package rediscache

import (
	"context"
	"github.com/caohui123/goweb/pkg/config"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"time"
)

var redisClient *redis.Client
var cacheClient *cache.Cache

// 初始化redisClient
func InitRedis(config config.Config) {
	redisCfg := config.RedisConfig
	redisClient = redis.NewClient(&redis.Options{
		DB:           redisCfg.Db,
		Addr:         redisCfg.Addr,
		Password:     redisCfg.Password,
		PoolSize:     redisCfg.PoolSize,
		MinIdleConns: redisCfg.MinIdleConns,
		IdleTimeout:  time.Duration(redisCfg.IdleTimeout) * time.Second,
	})
	_, err := redisClient.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}

	cacheClient = cache.New(&cache.Options{
		Redis:      redisClient,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
}

func GetRedisClient() *redis.Client {
	return redisClient
}
func GetCacheClient() *cache.Cache {
	return cacheClient
}

func CloseRedis() {
	_ = redisClient.Close()
}
