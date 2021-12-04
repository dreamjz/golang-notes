package v1

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	name := c.PostForm("name")
	pass := c.PostForm("password")
	if name != "kesa" || pass != ""
}