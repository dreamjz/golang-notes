package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/get", v1Get)
	}
	v2 := router.Group("/v2")
	{
		v2.GET("/get", v2Get)
	}
	router.Run(":9090")
}

func v1Get(c *gin.Context) {
	c.String(http.StatusOK, "V1 GET ")
}

func v2Get(c *gin.Context) {
	c.String(http.StatusOK, "V2 GET ")
}
