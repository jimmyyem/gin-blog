package v1

import (
	"gin-blog/controller"
	"gin-blog/models"
	"gin-blog/pkg/e"
	"gin-blog/pkg/logging"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

type TagController struct {
	controller.BaseController
}

// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [get]
func (t TagController) GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := models.CommonMaps()
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	code := e.SUCCESS
	data["lists"] = models.GetTags(util.GetOffset(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	res := t.ToJSON(code, data)
	c.JSON(http.StatusOK, res)
}

// @Summary Get multiple article tags
// @Produce  json
// @Param id path int false "Id"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tag/:id [get]
func (t TagController) GetTag(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))
	maps := models.CommonMaps()
	var data = models.Tag{}

	code := e.INVALID_PARAMS
	if id > 0 {
		code = e.SUCCESS
		maps["id"] = id
		data = models.GetTag(maps)
	}

	res := t.ToJSON(code, data)
	c.JSON(http.StatusOK, res)
}

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func (t TagController) AddTag(c *gin.Context) {
	name := c.PostForm("name")
	createdBy := c.PostForm("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if ! models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Error(err.Key + "=>" + err.Message)
		}
	}

	res := t.ToJSON(code, nil)
	c.JSON(http.StatusOK, res)
}

// @Summary 修改文章标签
// @Produce  json
// @Param id path int true "ID"
// @Param name query string true "ID"
// @Param state query int false "State"
// @Param modified_by query string true "ModifiedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags/{id} [put]
func (t TagController) EditTag(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))
	state := cast.ToInt(c.DefaultPostForm("state", "-1"))
	name := c.PostForm("name")
	modifiedBy := c.PostForm("modified_by")

	valid := validation.Validation{}
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Error(err.Key + "=>" + err.Message)
		}
	}

	res := t.ToJSON(code, nil)
	c.JSON(http.StatusOK, res)
}

// @Summary Delete article tag
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags/{id} [delete]
func (t TagController) DeleteTag(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Error(err.Key + "=>" + err.Message)
		}
	}

	res := t.ToJSON(code, nil)
	c.JSON(http.StatusOK, res)
}
