package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(MyGlobalMiddleware())

	router.GET("/req", func(c *gin.Context) {
		middleware := c.MustGet("middleware")
		c.String(http.StatusOK, fmt.Sprintf("Middleware type: %s", middleware))
	})

	router.GET("/req1", MyRouteMiddleware(), func(c *gin.Context) {
		middleware := c.MustGet("middleware")
		c.String(http.StatusOK, fmt.Sprintf("Middleware type: %s", middleware))
	})

	grp := router.Group("/grp")
	grp.Use(MyGroupMiddleware())
	{
		grp.GET("/req", func(c *gin.Context) {
			middleware := c.MustGet("middleware")
			c.String(http.StatusOK, fmt.Sprintf("Middleware type: %s", middleware))
		})
	}

	router.Run(":9090")
}

func MyGlobalMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("middleware", "global")
		fmt.Println("[Global] Before request")
		c.Next()
		fmt.Println("[Global] After request")
	}
}

func MyRouteMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("middleware", "route")
		fmt.Println("[Route] Before request")
		c.Next()
		fmt.Println("[Route] After request")
	}
}

func MyGroupMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("middleware", "group")
		fmt.Println("[Group] Before request")
		c.Next()
		fmt.Println("[Group] After request")
	}
}
