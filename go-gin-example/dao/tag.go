package dao

import (
	"errors"
	"go-gin-example/global"
	"go-gin-example/models"
	"go-gin-example/models/request"

	"gorm.io/gorm"
)

// GetTags get  tag list from database
func GetTags(offset int, pageSize int, condition models.Tag) (tags []models.Tag, err error) {
	err = global.AppDB.Model(&models.Tag{}).Where(&condition).Offset(offset).Limit(pageSize).Find(&tags).Error
	return
}

// GetTagTotal get counts of tags in database
func GetTagTotal(condition models.Tag) (count int64, err error) {
	err = global.AppDB.Model(&models.Tag{}).Where(&condition).Count(&count).Error
	return
}

// CreateTag insert a new tag to db
func CreateTag(tag models.Tag) (err error) {
	err = global.AppDB.Create(&tag).Error
	return
}

// EditTag edit tag info with specified data
func EditTag(id uint, tag request.EditTagReq) (err error) {
	data := map[string]interface{}{
		"name":       tag.Name,
		"state":      tag.State,
		"updated_by": tag.UpdatedBy,
	}
	err = global.AppDB.Model(&models.Tag{}).Where("id = ?", id).Updates(data).Error
	return
}

func DeleteTagByID(id uint) (err error) {
	err = global.AppDB.Delete(&models.Tag{}, id).Error
	return
}

// ExistTagByID return tag exists or not with specified id
func ExistTagByID(id uint) (bool, error) {
	var tag models.Tag
	err := global.AppDB.Model(&models.Tag{}).Where("id = ?", id).First(&tag).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return true, err
	}
	return true, nil
}

// ExistTagByName return same name tag exists or not
func ExistTagByName(name string) (bool, error) {
	var tag models.Tag
	err := global.AppDB.Model(&models.Tag{}).Where("name=?", name).First(&tag).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return true, err
	}
	return true, err
}
