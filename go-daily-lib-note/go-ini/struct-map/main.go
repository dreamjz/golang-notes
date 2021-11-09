package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
)

type Config struct {
	AppName  string `ini:"app_name"`
	LogLevel string `ini:"log_level"`

	MySql MysqlConfig `ini:"mysql"'`
	Redis RedisConfig `ini:"redis"`
}

type MysqlConfig struct {
	IP       string `ini:"ip"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type RedisConfig struct {
	IP   string `ini:"ip"`
	Port int    `ini:"port"`
}

func main() {
	config, err := ini.Load("../resources/config.ini")
	if err != nil {
		log.Fatal("load config error:", err)
	}
	c := Config{}
	config.MapTo(&c)
	fmt.Printf("%#v", c)
	// load from struct

	config2 := ini.Empty()
	c1 := Config{
		AppName:  "map_to_struct",
		LogLevel: "DEBUG",
		MySql: MysqlConfig{
			IP:   "127.0.0.1",
			Port: 3306,
			User: "root",
		},
		Redis: RedisConfig{
			IP: "127.0.0.1",
		},
	}

	err = ini.ReflectFrom(config2, &c1)
	if err != nil {
		log.Fatal("load from struct error")
	}
	err = config2.SaveTo("../resources/load_from_Struct.ini")
	if err != nil {
		log.Fatal("save config error", err)
	}

}
