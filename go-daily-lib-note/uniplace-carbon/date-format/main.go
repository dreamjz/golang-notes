package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05.000")) // 2021-11-06 03:57:13.845
	fmt.Println(now.Format("2006/01/02 15/04/05.000")) // 2021/11/06 03/57/13
	fmt.Println(now.Format("15:04:05.000"))            // 03:57:13.845
}
