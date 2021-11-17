package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/post", post)
	router.Run(":9090")
}

func post(c *gin.Context) {
	ids := c.QueryMap("ids")
	names := c.PostFormMap("names")
	c.JSON(http.StatusOK, gin.H{
		"ids":   ids,
		"names": names,
	})
}
