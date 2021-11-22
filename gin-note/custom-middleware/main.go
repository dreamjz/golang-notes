package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(myLogger())
	router.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example")
		log.Println(example)
	})
	log.Println("Listen and serve on 0.0.0.0:9090")
	router.Run(":9090")
}

func myLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "123456")
		c.Next()
		latency := time.Since(t)
		log.Println(latency)
		status := c.Writer.Status()
		log.Println(status)
	}
}
