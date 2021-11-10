package global

import (
	"go-gin-example/models/config"
	"gorm.io/gorm"
)

var (
	AppConfig config.Config
	AppDB     *gorm.DB
)
