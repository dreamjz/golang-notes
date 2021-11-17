package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/post", details)
	router.Run(":9090")
}

func details(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	message := c.PostForm("message")
	name := c.PostForm("name")
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"page":    page,
		"name":    name,
		"message": message,
	})
}
