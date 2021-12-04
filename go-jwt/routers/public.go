package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitializePublicRouter(router *gin.RouterGroup){
	publicRouter := router.Group("/public")
	{
		publicRouter.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK,"pong")
		})
	}
}
