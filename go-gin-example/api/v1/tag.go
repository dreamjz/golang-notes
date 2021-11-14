package v1

import (
	"errors"
	"fmt"
	"go-gin-example/global"
	"go-gin-example/models"
	"go-gin-example/models/request"
	"go-gin-example/models/response"
	"go-gin-example/service"
	"log"

	"github.com/spf13/cast"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// GetTags get tag list exists
func GetTags(c *gin.Context) {
	var tagReq request.QueryTagsPageReq
	err := c.ShouldBindQuery(&tagReq)
	if err != nil {
		log.Println(err)
		response.FailWithCode(response.InvalidParams, c)
		return
	}
	data, total, err := service.GetTags(tagReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(response.PageResult{
		Page:     tagReq.Page,
		PageSize: global.AppConfig.App.PageSize,
		List:     data,
		Total:    total,
	}, c)

}

// AddTag add a article tag
func AddTag(c *gin.Context) {
	var tag models.Tag
	err := c.ShouldBindJSON(&tag)
	if err != nil {
		log.Println(err)
		response.FailWithCode(response.InvalidParams, c)
		return
	}
	// validation
	valid := validation.Validation{}
	valid.Required(tag.Name, "name").Message("Name cannot be empty")
	valid.MaxSize(tag.Name, 100, "name").Message("Max length of name is 100")
	valid.Required(tag.CreatedBy, "created_by").Message("User created by cannot be empty")
	valid.MaxSize(tag.CreatedBy, 100, "createdBy").Message("Max length of created_by is 100")
	valid.Range(tag.State, 0, 1, "state").Message("State only be 0 or 1")
	if valid.HasErrors() {
		response.FailWithMessage(fmt.Sprint(valid.Errors), c)
		return
	}
	err = service.CreateTag(tag)
	if err != nil {
		if errors.Is(err, service.ErrorTagExists) {
			response.FailWithCode(response.ErrorExistTag, c)
			return
		}
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OK(c)
}

// EditTagById change tag by id
func EditTagById(c *gin.Context) {
	var tag request.EditTagReq
	err := c.ShouldBindJSON(&tag)
	log.Printf("Req:%#v", tag)
	if err != nil {
		log.Println(err)
		response.FailWithCode(response.InvalidParams, c)
		return
	}
	id, err := cast.ToUintE(c.Param("id"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// validation
	valid := validation.Validation{}
	valid.Required(tag.UpdatedBy, "updatedBy").Message("User updated by cannot be empty")
	valid.MaxSize(tag.UpdatedBy, 100, "updatedBy").Message("Max size of updatedBy is 100")
	valid.MaxSize(tag.Name, 100, "name").Message("Max size of name is 100")
	valid.Range(tag.State, 0, 1, "state").Message("State only be 0 or 1")
	if valid.HasErrors() {
		response.FailWithMessage(fmt.Sprint(valid.Errors), c)
		return
	}
	err = service.EditTag(id, tag)
	if errors.Is(err, service.ErrorTagNotExists) {
		response.FailWithCode(response.ErrorNotExistTag, c)
		return
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OK(c)
}

// DeleteTagById delete a tag by id
func DeleteTagById(c *gin.Context) {
	id, err := cast.ToUintE(c.Param("id"))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.DeleteTagByID(id)
	if errors.Is(err, service.ErrorTagNotExists) {
		response.FailWithCode(response.ErrorNotExistTag, c)
		return
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OK(c)
}
