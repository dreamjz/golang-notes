package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123"},
	"austin": gin.H{"email": "austin@bar.com", "phone": "456"},
	"lena":   gin.H{"email": "lena@bar.com", "phone": "789"},
}

func main() {
	router := gin.Default()
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"kim":    "4321",
	}))

	authorized.GET("/secrets", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "No Secret :( "})
	})
	log.Println("Listen and serve on 0.0.0.0:9090")
	router.Run(":9090")
}
