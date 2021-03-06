package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name       string    `form:"name"`
	Address    string    `form:"address"`
	Birthday   time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"8"`
	CreateTime time.Time `form:"createdTime" time_format:"unixNano"`
	UnixTime   time.Time `form:"unixTime" time_format:"unix"`
}

func main() {
	router := gin.Default()
	router.POST("/test", startPage)
	log.Println("Listen and serve on 0.0.0.0:9090")
	router.Run(":9090")
}

func startPage(c *gin.Context) {
	var p Person
	if err := c.ShouldBind(&p); err == nil {
		log.Println(p.Name)
		log.Println(p.Address)
		log.Println(p.Birthday)
		log.Println(p.CreateTime)
		log.Println(p.UnixTime)
	} else {
		c.String(http.StatusOK, err.Error())
		return
	}
	c.String(http.StatusOK, "Success")

}
