package routers

import (
	"github.com/gin-gonic/gin"
	v1 "go-gin-example/api/v1"
)

func InitArticleRouter(group *gin.RouterGroup) {
	articleGroup := group.Group("article")
	{
		articleGroup.GET("/:id", v1.GetArticleByID)
		articleGroup.GET("/articles", v1.GetArticles)
		articleGroup.POST("/create", v1.AddArticle)
		articleGroup.PUT("/edit/:id", v1.EditArticle)
		articleGroup.DELETE("/delete/:id", v1.DeleteArticleById)
	}
}
