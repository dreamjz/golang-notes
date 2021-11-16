package service

import (
	"errors"
	"go-gin-example/dao"
	"go-gin-example/utils"
	"log"

	"gorm.io/gorm"
)

var (
	ErrAuth = errors.New("auth error")
)

func CheckAuth(username, password string) (string, error) {
	err := dao.AuthByUserPass(username, password)
	var token string
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return token, ErrAuth
	}
	if err != nil {
		return token, err
	}
	token, err = utils.GenerateToken(username, password)
	if err != nil {
		log.Println("Generate token error")
		return token, err
	}
	return token, nil
}
