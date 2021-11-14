package main

import (
	"fmt"
	"go-gin-example/core"
	"go-gin-example/global"
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
	// Listen and serve
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", global.AppConfig.Server.HttpPort),
		Handler:        router,
		ReadTimeout:    global.AppConfig.Server.ReadTimeout,
		WriteTimeout:   global.AppConfig.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
