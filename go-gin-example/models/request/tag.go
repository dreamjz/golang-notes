package request

import (
	"go-gin-example/models"
)

type QueryTagsPageReq struct {
	models.Tag
	PageInfo
}

type EditTagReq struct {
	Name      string `json:"name"`
	State     int    `json:"state"`
	UpdatedBy string `json:"updatedBy"`
}
