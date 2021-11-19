package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin/binding"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

type Booking struct {
	CheckIn  time.Time `form:"checkin" binding:"required,bookableDate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"checkout" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func main() {
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookableDate", bookableDate)
	}
	router.GET("/bookable", getBookable)
	log.Println("Listen and serve on 0.0.0.0:9090")
	router.Run(":9090")
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Booking dates are valid",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}
