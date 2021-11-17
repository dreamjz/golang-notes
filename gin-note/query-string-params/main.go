package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/welcome", welcome)
	router.Run(":9090")
}

func welcome(c *gin.Context) {
	firstname := c.Query("firstname")
	lastname := c.Query("lastname")
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}
