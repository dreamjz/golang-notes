package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
)

func main() {
	config, err := ini.Load("../resources/config.ini")
	if err != nil {
		log.Fatal("read config error ", err)
	}
	fmt.Println("App name:", config.Section("").Key("app_name").String())
	fmt.Println("Log level:", config.Section("").Key("log_level").String())
	fmt.Println("Mysql ip:", config.Section("mysql").Key("ip").String())
	mysqlPort, err := config.Section("mysql").Key("port").Int()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mysqlPort)
	fmt.Println("Redis ip:", config.Section("redis").Key("ip").String())
	redisPort, err := config.Section("redis").Key("port").Int()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(redisPort)
}
