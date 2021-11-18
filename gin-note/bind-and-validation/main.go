package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/loginJSON", loginJSON)
	router.POST("/loginXML", loginXML)
	router.POST("/loginForm", loginForm)
	log.Println("Listen and serve on 0.0.0.0:8080")
	router.Run(":9090")
}

func loginJSON(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if json.User != "menu" || json.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "you are logged in",
	})
}

func loginXML(c *gin.Context) {
	var xml Login
	if err := c.ShouldBindXML(&xml); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if xml.User != "menu" || xml.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "you are logged in",
	})
}

func loginForm(c *gin.Context) {
	var form Login
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if form.User != "menu" || form.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "you are logged in",
	})
}
