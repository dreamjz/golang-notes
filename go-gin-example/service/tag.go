package service

import (
	"errors"
	"go-gin-example/dao"
	"go-gin-example/global"
	"go-gin-example/models"
	"go-gin-example/models/request"
	"go-gin-example/utils"
)

var (
	ErrorTagExists    = errors.New("tag exists")
	ErrorTagNotExists = errors.New("tag not exists")
)

// GetTags get  tag list from database
func GetTags(req request.QueryTagsPageReq) (tags []models.Tag, total int64, err error) {
	offset := utils.GetOffset(req.Page)
	pageSize := global.AppConfig.App.PageSize
	total, err = dao.GetTagTotal(req.Tag)
	if err != nil {
		return tags, total, err
	}
	tags, err = dao.GetTags(offset, pageSize, req.Tag)
	return tags, total, err
}

// CreateTag insert a new tag to db
func CreateTag(tag models.Tag) (err error) {
	exists, err := dao.ExistTagByName(tag.Name)
	if err != nil {
		return err
	}
	if exists {
		return ErrorTagExists
	}
	err = dao.CreateTag(tag)
	if err != nil {
		return err
	}
	return nil
}

// EditTag update tag info already exists
func EditTag(id uint, tag request.EditTagReq) (err error) {
	exists, err := dao.ExistTagByID(id)
	if err != nil {
		return err
	}
	if !exists {
		return ErrorTagNotExists
	}
	err = dao.EditTag(id, tag)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTagByID(id uint) (err error) {
	exists, err := dao.ExistTagByID(id)
	if err != nil {
		return err
	}
	if !exists {
		return ErrorTagNotExists
	}
	err = dao.DeleteTagByID(id)
	if err != nil {
		return err
	}
	return nil
}
