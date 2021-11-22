package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"html": "<h1>Hello</h1>",
		})
	})
	router.GET("/pureJson", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
			"html": "<h1>Hello</h1>",
		})
	})
	log.Println("Listen and serve on 0.0.0.0:9090")
	router.Run(":9090")
}
