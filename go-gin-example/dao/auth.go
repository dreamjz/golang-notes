package dao

import (
	"go-gin-example/global"
	"go-gin-example/models"
)

func AuthByUserPass(username, password string) (err error) {
	var auth models.Auth
	err = global.AppDB.Model(&models.Auth{}).Where("username = ? AND password = ?", username, password).First(&auth).Error
	return
}
