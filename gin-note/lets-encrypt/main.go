package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/autotls"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	log.Fatal(autotls.Run(router, "example.com", "example2.com"))
}
