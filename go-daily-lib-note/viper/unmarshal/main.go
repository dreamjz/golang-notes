package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	AppName  string `mapstructure:"app_name"`
	LogLevel string `mapstructure:"log_level"`

	MySql MySqlConfig
	Redis RedisConfig
}

type MySqlConfig struct {
	Port int
	ip   string // will not be parsed
}

type RedisConfig struct {
	Port int
}

func main() {
	viper.SetConfigName("config2")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../resources")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("read config failed:", err)
	}

	var c Config
	viper.Unmarshal(&c)

	fmt.Printf("%#v", c)
}
