package request

type PageInfo struct {
	Page int `form:"page" binding:"required"`
}
