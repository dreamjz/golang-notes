package models

type Article struct {
	Model
	TagID       uint   `json:"tagID" gorm:"index"`
	Tag         Tag    `json:"tag"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	State       int    `json:"state"`
	CreatedBy   string `json:"createdBy"`
	UpdatedBy   string `json:"updatedBy"`
}
