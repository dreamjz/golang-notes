package utils

import (
	"go-gin-example/global"
)

// GetOffset get page size from request and
// calculate offset for querying data in database
func GetOffset(page int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * global.AppConfig.App.PageSize
	}
	return result
}
