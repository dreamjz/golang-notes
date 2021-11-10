package core

import (
	"github.com/gin-gonic/gin"
	"go-gin-example/global"
	"go-gin-example/routers"
)

func Router() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gin.SetMode(global.AppConfig.RunMode)
	group := router.Group("/v1")
	routers.InitTagRouter(group)
	routers.InitArticleRouter(group)
	return router
}
