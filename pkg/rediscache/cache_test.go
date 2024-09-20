package rediscache

import (
	"context"
	"fmt"
	"github.com/caohui123/goweb/pkg/config"
	"github.com/go-redis/cache/v8"
	"os"
	"testing"
	"time"
)

type Object struct {
	Str string
	Num int
}

func init() {
	wd, _ := os.Getwd()
	conf := config.Load(wd + "/../../conf/config.yaml")
	InitRedis(*conf)
}
func Test_basicUsage(t *testing.T) {
	mycache := cache.New(&cache.Options{
		Redis:      GetRedisClient(),
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	ctx := context.TODO()
	key := "mykey"
	obj := &Object{
		Str: "mystring",
		Num: 42,
	}
	if !mycache.Exists(ctx, key) {
		fmt.Println("====================")
		if err := mycache.Set(&cache.Item{
			Ctx:   ctx,
			Key:   key,
			Value: obj,
			TTL:   time.Hour,
		}); err != nil {
			panic(err)
		}
	}
	var wanted Object
	if err := mycache.Get(ctx, key, &wanted); err == nil {
		fmt.Println(wanted)
	}
}
