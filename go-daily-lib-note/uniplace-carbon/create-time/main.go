package main

import (
	"fmt"
	"log"
	"time"

	"github.com/uniplaces/carbon"
)

func main() {
	// creat date with time
	fmt.Println("Create date with time:")
	loc, err := time.LoadLocation("Japan")
	if err != nil {
		log.Fatal("failed to load location:", err)
	}
	d := time.Date(2021, time.November, 4, 10, 51, 20, 0, loc)
	fmt.Printf("time in japan is :%s\n", d)
	// create date with uniplace-carbon
	fmt.Println("Create date with uniplace-carbon:")
	c, err := carbon.Create(2021, time.November, 4, 10, 51, 20, 0, "Japan")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("time in japan is :%s\n", c)
}
