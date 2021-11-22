package main

import (
	"log"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileForm struct {
	Name   string                `form:"name" binding:"required"`
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/profile", bindProfile)
	log.Println("Listen and serve on 0.0.0.0:9090")
	router.Run(":9090")
}

func bindProfile(c *gin.Context) {
	var form ProfileForm
	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusBadRequest, "bad request:", err.Error())
		return
	}
	err := c.SaveUploadedFile(form.Avatar, form.Avatar.Filename)
	if err != nil {
		c.String(http.StatusInternalServerError, "save file error:", err.Error())
		return
	}
	c.String(http.StatusOK, "success")
}
