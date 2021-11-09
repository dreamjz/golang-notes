package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
)

func main() {
	config, err := ini.Load("../resources/config.ini")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sections:%#v\n\n", config.Sections())
	fmt.Printf("Section names:%v\n\n", config.SectionStrings())

	newSection := config.Section("new_section")
	fmt.Printf("New section:%#v\n\n", newSection)
	fmt.Printf("Section names:%v\n\n", config.SectionStrings())

	_, err = config.NewSection("new")
	fmt.Println("Section names:", config.SectionStrings())
}
