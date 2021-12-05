package middleware

import (
	"go-jwt-note/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.ExtractToken(c)
		log.Println("Token: ", token)
		userClaims, err := utils.ParseToken(token)
		if err != nil {
			c.String(http.StatusForbidden, "unauthorized")
			c.Abort()
			return
		}
		err = utils.FetchAuth(userClaims.UUID)
		if err != nil {
			c.String(http.StatusForbidden, "unauthorized")
			c.Abort()
			return
		}
		c.Set("userClaims", userClaims)
		c.Next()
	}
}
