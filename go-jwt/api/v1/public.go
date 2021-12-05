package v1

import (
	"go-jwt-note/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserTokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func Login(c *gin.Context) {
	name := c.PostForm("username")
	pass := c.PostForm("password")
	if name != "kesa" || pass != "123" {
		c.String(http.StatusUnauthorized, "please provide valid login details")
		return
	}
	td, err := utils.GenerateToken(name)
	if err != nil {
		c.String(http.StatusUnprocessableEntity, err.Error())
		return
	}
	err = utils.SaveUserTokens(name, td)
	if err != nil {
		c.String(http.StatusUnprocessableEntity, err.Error())
		return
	}
	tokens := UserTokens{
		AccessToken:  td.AccessToken,
		RefreshToken: td.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
}

func RefreshToken(c *gin.Context) {
	token := utils.ExtractToken(c)
	userClaims, err := utils.ParseToken(token)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		return
	}
	err = utils.FetchAuth(userClaims.UUID)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		return
	}
	// refresh token valid
	// delete previous refresh token
	name := userClaims.Username
	deleted, err := utils.RemoveAuth(userClaims.UUID)
	if err != nil || deleted == 0 {
		c.String(http.StatusUnauthorized, err.Error())
		return
	}
	// generate new access token and refresh token
	td, err := utils.GenerateToken(name)
	if err != nil {
		c.String(http.StatusUnprocessableEntity, err.Error())
		return
	}
	err = utils.SaveUserTokens(name, td)
	if err != nil {
		c.String(http.StatusUnprocessableEntity, err.Error())
		return
	}
	tokens := UserTokens{
		AccessToken:  td.AccessToken,
		RefreshToken: td.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
}
