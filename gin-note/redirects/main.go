package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/bing", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.bing.com")
	})
	router.GET("/welcome", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome")
	})
	router.POST("/login", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/welcome")
	})
	router.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		router.HandleContext(c)
	})
	router.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})
	router.Run(":9090")
}
