package v1

import (
	"go-jwt-note/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	userClaims, _ := c.MustGet("userClaims").(*utils.UserClaims)
	c.String(http.StatusOK, "Welcome ! %s", userClaims.Username)
}

func Logout(c *gin.Context) {
	userClaims, _ := c.MustGet("userClaims").(*utils.UserClaims)
	deleted, err := utils.RemoveAuth(userClaims.UUID)
	if err != nil || deleted == 0 {
		c.String(http.StatusUnauthorized, "unauthorized")
		return
	}
	c.String(http.StatusOK, "Logout success")
}
