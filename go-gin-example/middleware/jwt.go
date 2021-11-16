package middleware

import (
	"go-gin-example/models/response"
	"go-gin-example/utils"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			log.Println("token is empty")
			response.FailWithCode(response.InvalidParams, c)
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			log.Println("parse token error")
			response.FailWithCode(response.ErrorAuthCheckTokenFail, c)
			c.Abort()
			return
		}
		if time.Now().Unix() > claims.ExpiresAt {
			response.FailWithCode(response.ErrorAuthCheckTokenTimeout, c)
			c.Abort()
			return
		}
		c.Next()
	}
}
