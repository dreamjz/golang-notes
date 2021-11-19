package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin/binding"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `binding:"required,email"`
}

func UserStructLevelValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(User)
	if len(user.Firstname) == 0 && len(user.Lastname) == 0 {
		sl.ReportError(user.Firstname, "Firstname", "Firstname", "FirstnameOrLastname", "")
		sl.ReportError(user.Lastname, "Lastname", "Lastname", "FirstnameOrLastname", "")
	}

}

func main() {
	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(UserStructLevelValidation, User{})
	}
	router.POST("/user", validateUser)
	log.Println("Listen and serve on 0.0.0.0:9090")
	router.Run(":9090")
}

func validateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "User validation successful",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User validation failed",
			"error":   err.Error(),
		})
	}
}
