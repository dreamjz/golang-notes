package utils

import (
	"context"
	"errors"
	"go-jwt-note/global"
	"log"
	"time"
)

var ctx = context.Background()

var (
	ErrSaveToken   = errors.New("save token error")
	ErrFetchToken  = errors.New("fetch token error")
	ErrRemoveToken = errors.New("remove token error")
)

func SaveUserTokens(username string, td *tokenDetails) error {
	rdb := global.RedisDB
	now := time.Now()
	atExp := td.AtExpires.Sub(now)
	rtExp := td.RtExpires.Sub(now)

	err := rdb.Set(ctx, td.AccessUUID, username, atExp).Err()
	if err != nil {
		log.Println("Save access token failed: ", err.Error())
		return ErrSaveToken
	}
	err = rdb.Set(ctx, td.RefreshUUID, username, rtExp).Err()
	if err != nil {
		log.Println("Save refresh token failed: ", err.Error())
		return ErrSaveToken
	}
	return nil
}

func FetchAuth(tokenUUID string) error {
	_, err := global.RedisDB.Get(ctx, tokenUUID).Result()
	if err != nil {
		log.Println("Get token failed: ", err.Error())
		return ErrFetchToken
	}
	return nil
}

func RemoveAuth(tokenUUID string) (int64, error) {
	deleted, err := global.RedisDB.Del(ctx, tokenUUID).Result()
	if err != nil {
		log.Println("Remove token failed: ", err.Error())
		return 0, ErrRemoveToken
	}
	log.Printf("%d keys deleted", deleted)
	return deleted, nil
}
