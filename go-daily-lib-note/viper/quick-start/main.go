package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("../resources")
	// set default value
	viper.SetDefault("redis.port", 6379)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: ", err)
	}

	// print config
	// base config
	fmt.Println("app name", viper.Get("app_name"))
	fmt.Println("log level", viper.Get("log_level"))

	// mysql config
	fmt.Println("mysql ip:", viper.Get("mysql.ip"))
	fmt.Println("mysql port:", viper.Get("mysql.port"))
	fmt.Println("mysql user:", viper.Get("mysql.user"))
	fmt.Println("mysql pass:", viper.Get("mysql.password"))
	fmt.Println("mysql database:", viper.Get("mysql.database"))

	// redis config
	fmt.Println("redis ip:", viper.Get("redis.ip"))
	fmt.Println("redis port:", viper.Get("redis.port"))

	pflag.Args()
}
