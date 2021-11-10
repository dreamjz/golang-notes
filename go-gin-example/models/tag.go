package models

import "go-gin-example/core"

type Tag struct {
	core.Model
	Name      string `json:"name"`
	State     int    `json:"state"`
	CreatedBy string `json:"createdBy"`
	UpdatedBy string `json:"updatedBy"`
}
