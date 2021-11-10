package routers

import (
	"github.com/gin-gonic/gin"
	v1 "go-gin-example/api/v1"
)

func InitTagRouter(group *gin.RouterGroup) {
	tagGroup := group.Group("article")
	{
		tagGroup.GET("/tags", v1.GetTags)
		tagGroup.POST("/tags", v1.AddTag)
		tagGroup.PUT("/tags", v1.EditTagById)
		tagGroup.DELETE("/tags", v1.DeleteTagById)
	}
}
