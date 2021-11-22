package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

func main() {
	router := gin.Default()
	router.POST("/bindCheckboxes", bindCheckboxes)
	log.Println("Listen and serve on 0.0.0.0:9090")
	router.Run(":9090")
}

func bindCheckboxes(c *gin.Context) {
	var f myForm
	if err := c.ShouldBind(&f); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, f)
}
