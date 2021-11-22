package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./serving-static-files/assets")
	router.StaticFS("/more_static", http.Dir("./serving-static-files/more_static"))
	router.StaticFS("/more_static2", gin.Dir("./serving-static-files/more_static", false))
	router.StaticFile("/favicon.ico", "./serving-static-files/resources/favicon.ico")
	log.Println("Listen and serve on 0.0.0.0:9090")
	router.Run(":9090")
}
