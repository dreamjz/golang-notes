package routers

import "github.com/gin-gonic/gin"

func InitArticleRouter(group *gin.RouterGroup) {
	tagGroup := group.Group("tag")
	{
		tagGroup.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "test",
			})
		})
	}
}
