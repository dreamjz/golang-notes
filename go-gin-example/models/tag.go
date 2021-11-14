package models

type Tag struct {
	Model
	Name      string `json:"name" form:"name""`
	State     int    `json:"state" form:"state"`
	CreatedBy string `json:"createdBy" form:"createdBy"`
	UpdatedBy string `json:"updatedBy" form:"updatedBy"`
}
