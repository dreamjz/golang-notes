package utils

import (
	"go-gin-example/global"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	now := time.Now()
	expiredTime := now.Add(3 * time.Hour)

	claims := Claims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Println(global.AppConfig.App.JwtSecret)
	token, err := tokenClaims.SignedString([]byte(global.AppConfig.App.JwtSecret))
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.AppConfig.App.JwtSecret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
