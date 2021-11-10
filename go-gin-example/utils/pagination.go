package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-gin-example/global"
)

// GetPage get page size from request and
// calculate offset for querying data in database
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * global.AppConfig.App.PageSize
	}
	return result
}
