package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/favicon", func(c *gin.Context) {
		c.File("./serving-data-from-file/resources/favicon.ico")
	})
	fs := http.FileSystem(http.Dir("././serving-data-from-file/assets"))
	router.GET("/pacman", func(c *gin.Context) {
		c.FileFromFS("Pac Man.ico", fs)
	})
	log.Println("Listen and serve on 0.0.0.0:9090")
	router.Run(":9090")
}
