package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/long_async", func(c *gin.Context) {
		cCp := c.Copy()
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})
	router.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("Done! in path " + c.Request.URL.Path)
	})
	router.Run(":9090")
}
