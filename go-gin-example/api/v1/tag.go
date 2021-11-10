package v1

import (
	"github.com/gin-gonic/gin"
	"go-gin-example/service"
)

// GetTags get tag list exists
func GetTags(c *gin.Context) {
	data := service.GetTags(0,5,map[string]interface{}{"name":"s"})
	c.JSON(200,data)
}

// AddTag add a article tag
func AddTag(c *gin.Context) {

}

// EditTagById change tag by id
func EditTagById(c *gin.Context) {

}

// DeleteTagById delete a tag by id
func DeleteTagById(c *gin.Context) {

}
