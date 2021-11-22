package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ASCII-JSON", func(c *gin.Context) {
		c.AsciiJSON(http.StatusOK, gin.H{
			"lang": "GO语言",
			"tag":  "<br>",
		})
	})
	log.Println("Listen and serve on 0.0.0.0:9090")
	router.Run(":9090")
}
