package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/get", getting)
	router.POST("/post", posting)
	router.PUT("/put", putting)
	router.DELETE("/delete", deleting)
	router.PATCH("/patch", patching)
	router.HEAD("/head", head)
	router.OPTIONS("/options", options)

	router.Run(":9090")
}

func getting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GET",
	})
}

func posting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "POST",
	})
}

func putting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "PUT",
	})
}

func deleting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "DELETE",
	})
}

func patching(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func head(c *gin.Context) {
	c.Header("key", "val")
}

func options(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OPTION",
	})
}
