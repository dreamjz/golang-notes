package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

func main() {
	config := ini.Empty()

	defaultSection := config.Section("")
	defaultSection.NewKey("app_name", "save config")
	defaultSection.NewKey("log_level", "DEBUG")

	mysqlSec, err := config.NewSection("mysql")
	if err != nil {
		fmt.Println("create mysql section error:", err)
	}
	mysqlSec.NewKey("ip", "127.0.0.1")
	mysqlSec.NewKey("port", "3306")

	redisSec, err := config.NewSection("reids")
	if err != nil {
		fmt.Println("create redis section error :", err)
	}
	redisSec.NewKey("ip", "127.0.0.1")
	redisSec.NewKey("port", "6379")

	err = config.SaveTo("../resources/saved-config.ini")
	if err != nil {
		fmt.Println("save config error :", err)
	}

	err = config.SaveToIndent("../resources/saved-config-pretty.ini", " ")
	if err != nil {
		fmt.Println("save config error:", err)
	}

	config.WriteTo(os.Stdout)
	fmt.Println()
	config.WriteToIndent(os.Stdout, " ")

}
