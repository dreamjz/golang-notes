package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()

	viper.BindEnv("redis.port", "redis_port_1", "redis_port_2", "redis_port3")
}

func main() {
	fmt.Println("HOME:", viper.GetString("HOME"))
	fmt.Println("redis.port", viper.GetInt("redis.port"))
}
