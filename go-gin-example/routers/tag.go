package routers

import (
	v1 "go-gin-example/api/v1"

	"github.com/gin-gonic/gin"
)

func InitTagRouter(group *gin.RouterGroup) {
	tagGroup := group.Group("tag")
	{
		tagGroup.GET("/tags", v1.GetTags)
		tagGroup.POST("/create", v1.AddTag)
		tagGroup.PUT("/edit/:id", v1.EditTagById)
		tagGroup.DELETE("/delete/:id", v1.DeleteTagById)
	}
}
