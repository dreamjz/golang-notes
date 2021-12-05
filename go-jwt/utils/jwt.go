package utils

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"

	"github.com/golang-jwt/jwt/v4"
)

const (
	secret = "my-secret"
)

var (
	ErrGenerateToken = errors.New("generate token error")
	ErrParseToken    = errors.New("parse token error")
)

type UserClaims struct {
	Username  string `json:"username"`
	UUID      string `json:"uuid"`
	TokenType string `json:"type"`
	jwt.RegisteredClaims
}

type tokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	// access token expires time
	AtExpires time.Time
	// refresh token expires time
	RtExpires time.Time
}

func GenerateToken(username string) (*tokenDetails, error) {
	rtExpiresTime := time.Now().Add(5 * time.Minute)
	atExpireTime := time.Now().Add(1 * time.Minute)
	atUUID := uuid.New().String()
	rtUUID := uuid.New().String()
	accessToken, err := createToken(username, "access", atUUID, atExpireTime)
	if err != nil {
		return nil, ErrGenerateToken
	}
	refreshToken, err := createToken(username, "refresh", rtUUID, rtExpiresTime)
	if err != nil {
		return nil, ErrGenerateToken
	}
	td := &tokenDetails{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		AccessUUID:   atUUID,
		RefreshUUID:  rtUUID,
		AtExpires:    atExpireTime,
		RtExpires:    rtExpiresTime,
	}
	return td, nil
}

func createToken(username, tokenType, tokenUUID string, expiresTime time.Time) (string, error) {
	claims := UserClaims{
		Username:  username,
		UUID:      tokenUUID,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "jwt-demo",
			ExpiresAt: jwt.NewNumericDate(expiresTime),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	if err != nil {
		log.Printf("Generate token for user: %s, err: %s", username, err.Error())
		return "", err
	}
	return token, nil
}

func ParseToken(token string) (*UserClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		log.Println("Parse token failed: ", err.Error())
		return nil, ErrParseToken
	}

	claims, ok := tokenClaims.Claims.(*UserClaims)
	if !ok {
		return nil, ErrParseToken
	}

	return claims, nil
}

func ExtractToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	log.Println("Authorization: ", authHeader)
	token := strings.TrimPrefix(authHeader, "Bearer ")
	return token
}
