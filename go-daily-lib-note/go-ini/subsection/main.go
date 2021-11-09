package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
)

func main() {
	config, err := ini.Load("../resources/sub_section.ini")
	if err != nil {
		log.Fatal("Read config error:", err)
	}
	fmt.Println("CLONE_URL:", config.Section("package.sub").Key("CLONE_URL").String())
}
