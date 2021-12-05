package initialize

import (
	"go-jwt-note/middleware"
	"go-jwt-note/routers"
	"log"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	// public router
	publicGroup := router.Group("")
	routers.InitializePublicGroup(publicGroup)
	// private router, needs authorization
	privateGroup := router.Group("")
	privateGroup.Use(middleware.Jwt())
	routers.InitializeUserRouter(privateGroup)
	server := endless.NewServer(":9090", router)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("Start server failed: %s", err.Error())
	}
}
