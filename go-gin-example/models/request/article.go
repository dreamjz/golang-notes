package request

import "go-gin-example/models"

type QueryArticlePageReq struct {
	models.Article
	PageInfo
}
