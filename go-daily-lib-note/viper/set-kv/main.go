package main

import (
	"fmt"
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	// default value of redis.port is 6379
	//pflag.Int("redis.port", 6379, "redis port")

	// 绑定命令行
	viper.BindPFlags(pflag.CommandLine)

}

func main() {
	pflag.Parse()

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("../resources")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("read config failed :", err)
	}

	viper.Set("mysql.port", 3307)

	fmt.Println("Mysql port:", viper.GetInt("mysql.port"))
	// redis.port = 7381 in config.toml
	fmt.Println("Redis port", viper.GetInt("redis.port"))

}
