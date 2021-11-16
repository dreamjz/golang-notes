package routers

import (
	v1 "go-gin-example/api/v1"

	"github.com/gin-gonic/gin"
)

func InitAuthRouter(group *gin.RouterGroup) {
	authGroup := group.Group("/auth")
	{
		authGroup.POST("/login", v1.Login)
	}
}
