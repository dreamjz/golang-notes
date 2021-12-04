package routers

import "github.com/gin-gonic/gin"

func InitializeUserRouter(router *gin.RouterGroup){
	userRouter := router.Group("/user")
	{
		router.GET("/welcome", func(context *gin.Context) {

		})
	}
}
