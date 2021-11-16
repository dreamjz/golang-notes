package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go-gin-example/global"
	"go-gin-example/models"
	"go-gin-example/models/request"
	"go-gin-example/models/response"
	"go-gin-example/service"
)

func GetArticleByID(c *gin.Context) {
	id, err := cast.ToUintE(c.Param("id"))
	if err != nil {
		response.FailWithCode(response.InvalidParams, c)
		return
	}
	data, err := service.GetArticleByID(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(data, c)
}

func GetArticles(c *gin.Context) {
	var req request.QueryArticlePageReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	articles, total, err := service.GetArticles(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(response.PageResult{
		List:     articles,
		Total:    total,
		Page:     req.Page,
		PageSize: global.AppConfig.App.PageSize,
	}, c)

}

func AddArticle(c *gin.Context) {
	var req models.Article
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.AddArticle(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OK(c)
}

func EditArticle(c *gin.Context) {
	id, err := cast.ToUintE(c.Param("id"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var req request.EditArticleReq
	err = c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.EditArticle(id, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OK(c)
}

func DeleteArticleById(c *gin.Context) {
	id, err := cast.ToUintE(c.Param("id"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.DeleteArticleByID(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OK(c)
}
