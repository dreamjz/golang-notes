package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/form_post", formPost)
	router.Run(":9090")
}

func formPost(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")
	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}
