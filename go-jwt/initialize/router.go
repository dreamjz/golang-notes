package initialize

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"go-jwt-note/routers"
)


func Run() {
	router := gin.Default()
	pulicGroup := router.Group("")
	{
		routers.InitializePublicRouter(pulicGroup)
	}
	privateGroup := router.Group("/v1")
	{

	}
	server := endless.NewServer(":9090",router)
	server.ListenAndServe()
}
