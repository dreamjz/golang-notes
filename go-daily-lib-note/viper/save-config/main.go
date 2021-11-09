package main

import (
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigName("config3")
	viper.SetConfigType("json")
	viper.AddConfigPath("../resources")

	viper.Set("app_name", "save-config")
	viper.Set("mysql.port", 3306)
	viper.Set("redis.port", 6379)

	if err := viper.SafeWriteConfig(); err != nil {
		log.Fatal("write config fialed :", err)
	}
}
