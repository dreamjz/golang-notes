package v1

import (
	"errors"
	"go-gin-example/models/response"
	"go-gin-example/service"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token, err := service.CheckAuth(username, password)
	if errors.Is(err, service.ErrAuth) {
		response.FailWithCode(response.ErrorAuthToken, c)
		return
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(token, c)
}
