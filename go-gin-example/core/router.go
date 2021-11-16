package core

import (
	"go-gin-example/global"
	"go-gin-example/middleware"
	"go-gin-example/routers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gin.SetMode(global.AppConfig.RunMode)
	publicGroup := router.Group("/")
	routers.InitAuthRouter(publicGroup)
	group := router.Group("/v1")
	group.Use(middleware.JWT())
	routers.InitTagRouter(group)
	routers.InitArticleRouter(group)
	return router
}
