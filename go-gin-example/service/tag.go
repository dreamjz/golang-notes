package service

import (
	"go-gin-example/global"
	"go-gin-example/models"
)

func GetTags(pageNum int,pageSize int,maps map[string]interface{}) (tags []models.Tag) {
	global.AppDB.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}