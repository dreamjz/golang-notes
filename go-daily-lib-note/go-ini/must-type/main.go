package main

import (
	"fmt"
	"gopkg.in/ini.v1"
)

func main() {
	config, err := ini.Load("../resources/config_2.ini")
	if err != nil {
		fmt.Println("read config error :", err)
	}
	fmt.Println("Before call MustType")
	port, err := config.Section("redis").Key("port").Int()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("port:", port)
	fmt.Println("Call MustType")
	port = config.Section("redis").Key("port").MustInt(6379)
	fmt.Println("port:", port)
	port, err = config.Section("redis").Key("port").Int()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("After call MustType")
	fmt.Println("port:", port)
}
