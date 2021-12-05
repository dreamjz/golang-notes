package routers

import (
	v1 "go-jwt-note/api/v1"

	"github.com/gin-gonic/gin"
)

func InitializeUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("/user")
	{
		userRouter.GET("/welcome", v1.Welcome)
		userRouter.GET("/logout", v1.Logout)
	}
}
