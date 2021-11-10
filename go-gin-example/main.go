package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-example/core"
	"go-gin-example/global"
)

func main() {
	core.Viper()
	core.Gorm()
	core.CloseDB()
	fmt.Printf("%#v\n", global.AppConfig)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":9090")
}
