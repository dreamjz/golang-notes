package routers

import (
	v1 "go-jwt-note/api/v1"

	"github.com/gin-gonic/gin"
)

func InitializePublicGroup(router *gin.RouterGroup) {
	publicRouter := router.Group("/public")
	{
		publicRouter.GET("/ping", v1.Ping)
		publicRouter.POST("/login", v1.Login)
		publicRouter.GET("/refresh", v1.RefreshToken)
	}
}
