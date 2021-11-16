package request

import "go-gin-example/models"

type QueryArticlePageReq struct {
	models.Article
	PageInfo
}

type EditArticleReq struct {
	TagID       uint   `json:"tagID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	State       int    `json:"state"`
	UpdatedBy   string `json:"updatedBy"`
}
