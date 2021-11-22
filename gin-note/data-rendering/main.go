package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin/testdata/protoexample"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/someJSON", someJSON)
	router.GET("/moreJSON", moreJSON)
	router.GET("/someXML", someXML)
	router.GET("/someYAML", someYAML)
	router.GET("/someProtoBuf", someProtoBuf)
	log.Println("Listen and serve on 0.0.0.0:9090")
	router.Run(":9090")
}

func someJSON(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hey",
		"status":  http.StatusOK,
	})
}

func moreJSON(c *gin.Context) {
	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg.Name = "Kesa"
	msg.Message = "hey"
	msg.Number = 123
	c.JSON(http.StatusOK, msg)
}

func someXML(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{
		"message": "hey",
		"status":  http.StatusOK,
	})
}

func someYAML(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{
		"message": "hey",
		"status":  http.StatusOK,
	})
}

func someProtoBuf(c *gin.Context) {
	reps := []int64{int64(1), int64(2)}
	label := "test"
	data := &protoexample.Test{
		Label: &label,
		Reps:  reps,
	}
	c.ProtoBuf(http.StatusOK, data)
}
