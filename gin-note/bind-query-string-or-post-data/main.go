package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	router := gin.Default()
	router.Any("/test", startPage)
	log.Println("Listen and serve on 0.0.0.0:9090")
	router.Run(":9090")
}

func startPage(c *gin.Context) {
	var p Person
	if c.ShouldBindQuery(&p) == nil {
		log.Println("===== Only Bind Query String =====")
		log.Println(p.Name)
		log.Println(p.Address)
	}
	c.String(http.StatusOK, "Success")

}
