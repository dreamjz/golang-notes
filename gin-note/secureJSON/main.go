package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/secureJSON", func(c *gin.Context) {
		names := []string{"lena", "kesa", "foo"}
		c.SecureJSON(http.StatusOK, names)
	})
	log.Println("Listen and serve on 0.0.0.0:9090")
	router.Run(":9090")
}
