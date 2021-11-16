package dao

import (
	"errors"
	"go-gin-example/global"
	"go-gin-example/models"

	"gorm.io/gorm"
)

func CreateArticle(data models.Article) (err error) {
	err = global.AppDB.Create(&data).Error
	return
}
func UpdateArticle(id uint, data models.Article) (err error) {
	err = global.AppDB.Model(&models.Article{}).Where("id = ?", id).Updates(data).Error
	return
}

func GetArticles(offset, limit int, cond models.Article) (articles []models.Article, err error) {
	err = global.AppDB.Preload("Tag").Model(&models.Article{}).Offset(offset).Limit(limit).Where(&cond).Find(&articles).Error
	return
}

func GetArticleByID(id uint) (article models.Article, err error) {
	err = global.AppDB.Preload("Tag").Model(&models.Article{}).Where("id = ?", id).First(&article).Error
	return
}

func DeleteArticleByID(id uint) (err error) {
	err = global.AppDB.Delete(&models.Article{}, id).Error
	return
}

func GetArticleTotal(cond models.Article) (count int64, err error) {
	err = global.AppDB.Model(&models.Article{}).Where(cond).Count(&count).Error
	return
}

func ExistsArticleByID(id uint) (bool, error) {
	var article models.Article
	err := global.AppDB.Model(&models.Article{}).Where("id =?", id).First(&article).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return true, err
	}
	return true, nil
}
