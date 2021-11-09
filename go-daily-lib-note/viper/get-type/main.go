package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("../resources")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("read config failed:", err)
	}

	fmt.Println("protocols:", viper.GetStringSlice("server.protocols"))
	fmt.Println("ports:", viper.GetIntSlice("server.ports"))
	fmt.Println("timeout:", viper.GetDuration("server.timeout"))

	fmt.Println("mysql ip:", viper.GetString("mysql.ip"))
	fmt.Println("mysql port", viper.GetInt("mysql.port"))

	if viper.IsSet("redis.port") {
		fmt.Println("redis port is set ")
	} else {
		fmt.Println("redis port is not set")
	}

	fmt.Println("mysql settings: ", viper.GetStringMap("redis"))
	fmt.Println("mysql settings: ", viper.GetStringMapString("mysql"))
	fmt.Println("all settings:", viper.AllSettings())
}
