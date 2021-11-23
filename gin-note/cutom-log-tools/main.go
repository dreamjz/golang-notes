package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	router.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "get")
	})
	router.POST("/post", func(c *gin.Context) {
		c.String(http.StatusOK, "post")
	})
	router.Run(":9090")
}
