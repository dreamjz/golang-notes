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
	ErrArticleNotExists = errors.New("article not exists")
)

func GetArticleByID(id uint) (models.Article, error) {
	return dao.GetArticleByID(id)
}

func GetArticles(req request.QueryArticlePageReq) ([]models.Article, int64, error) {
	offset := utils.GetOffset(req.Page)
	limit := global.AppConfig.App.PageSize
	total, err := dao.GetArticleTotal(req.Article)
	if err != nil {
		return nil, 0, err
	}
	articles, err := dao.GetArticles(offset, limit, req.Article)
	return articles, total, err
}

func AddArticle(req models.Article) error {
	return dao.CreateArticle(req)
}

func DeleteArticleByID(id uint) error {
	exists, err := dao.ExistsArticleByID(id)
	if err != nil {
		return err
	}
	if exists {
		return dao.DeleteArticleByID(id)
	}
	return ErrArticleNotExists
}

func EditArticle(id uint, req models.Article) error {
	return dao.UpdateArticle(id, req)
}
