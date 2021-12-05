package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		log.Println("set error: ", err.Error())
	}
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		log.Println("get error: ", err.Error())
	}
	fmt.Println("key:", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		log.Println("key2 does not exist")
	} else if err != nil {
		log.Println("get error: ", err.Error())
	} else {
		fmt.Println("key2:", val2)
	}
}
