package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/JSONP", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{
			"foo": "bar",
		})
	})
	log.Println("Listen and serve on 0.0.0.0:9090")
	router.Run(":9090")
}
