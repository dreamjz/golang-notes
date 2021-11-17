package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/user/:name", hello)
	router.GET("/user/:name/*action", action)
	router.POST("/user/:name/*action", fullPath)
	router.GET("/user/groups", groups)

	router.Run(":9090")
}

func hello(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func action(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}

func fullPath(c *gin.Context) {
	path := c.FullPath()
	b := path == "/user/:name/*action"
	c.String(http.StatusOK, "path: %s,%t", path, b)
}

func groups(c *gin.Context) {
	c.String(http.StatusOK, "The available groups are [...]")
}
