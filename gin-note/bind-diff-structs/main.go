package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
)

type FormA struct {
	Foo string `json:"foo" binding:"required"`
}

type FormB struct {
	Bar string `json:"bar" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/test", func(c *gin.Context) {
		objA := FormA{}
		objB := FormB{}
		if errA := c.ShouldBind(&objA); errA == nil {
			c.String(http.StatusOK, "the body should be fromA")
		}
		if errB := c.ShouldBind(&objB); errB == nil {
			c.String(http.StatusOK, "the body should be fromA")
		} else {
			log.Println(errB)
		}

	})
	router.POST("/test2", func(c *gin.Context) {
		objA := FormA{}
		objB := FormB{}
		if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
			c.String(http.StatusOK, "the body should be fromA")
		}
		if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
			c.String(http.StatusOK, "the body should be fromB")
		}
	})
	router.Run(":9090")
}
