package initialize

import (
	"context"
	"go-jwt-note/global"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func Redis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "my-redis:6379",
		Password: "",
		DB:       0,
	})
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatalln("failed to connect redis: ", err.Error())
	}
	log.Println("Connect redis success")
	global.RedisDB = rdb
}
