package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"time"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("../resources")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: ", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file:%s Op:%s\n", e.Name, e.Op)
	})

	fmt.Println("Before : redis.port=", viper.GetInt("redis.port"))
	time.Sleep(time.Second * 10)
	fmt.Println("After : redis.port=", viper.GetInt("redis.port"))

}
