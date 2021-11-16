package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-example/core"
	"go-gin-example/global"
	"log"
	"net/http"
)

func main() {
	// initialize
	core.Viper()
	core.Gorm()
	if global.AppDB != nil {
		db, _ := global.AppDB.DB()
		defer db.Close()
	}
	// router
	router := core.Router()
	// test
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	// Listen and serve
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", global.AppConfig.Server.HttpPort),
		Handler:        router,
		ReadTimeout:    global.AppConfig.Server.ReadTimeout,
		WriteTimeout:   global.AppConfig.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Serving at port : 9090")
	s.ListenAndServe()

}
