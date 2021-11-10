package v1

import (
	"fmt"
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

type ArticleControoler struct {
	controller.BaseController
}

// @Summary Get a single article
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles/{id} [get]
func (a ArticleControoler) GetArticle(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if ! valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		errMap := make(map[string]string)
		for _, err := range valid.Errors {
			//logging.Info(err.Key, err.Message)
			errMap[err.Key] = err.Message
		}
		fmt.Println(errMap)
	}

	// 构造返回内容
	res := a.ToJSON(code, data)

	c.JSON(http.StatusOK, res)
}

// @Summary Get multiple articles
// @Produce  json
// @Param tag_id body int false "TagID"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles [get]
func (a ArticleControoler) GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := models.CommonMaps()

	valid := validation.Validation{}
	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = cast.ToInt(arg)
		maps["tag_id"] = tagId
		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		data["lists"] = models.GetArticles(util.GetOffset(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	} else {
		for _, err := range valid.Errors {
			logging.Error(err.Key + "=>" + err.Message)
		}
	}

	res := a.ToJSON(code, data)
	c.JSON(http.StatusOK, res)
}

// @Summary Add article
// @Produce  json
// @Param tag_id body int true "TagID"
// @Param title body string true "Title"
// @Param desc body string true "Desc"
// @Param content body string true "Content"
// @Param created_by body string true "CreatedBy"
// @Param state body int true "State"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles [post]
func (a ArticleControoler) AddArticle(c *gin.Context) {
	tagId := cast.ToInt(c.PostForm("tag_id"))
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	content := c.PostForm("content")
	createdBy := c.PostForm("created_by")

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistTagByID(tagId) {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = models.STATE_ONLINE
			models.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Error(err.Key + "=>" + err.Message)
		}
	}

	res := a.ToJSON(code, nil)
	c.JSON(http.StatusOK, res)
}

// @Summary Update article
// @Produce  json
// @Param id path int true "ID"
// @Param tag_id body string false "TagID"
// @Param title body string false "Title"
// @Param desc body string false "Desc"
// @Param content body string false "Content"
// @Param modified_by body string true "ModifiedBy"
// @Param state body int false "State"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles/{id} [put]
func (a ArticleControoler) EditArticle(c *gin.Context) {
	valid := validation.Validation{}
	id := cast.ToInt(c.Param("id"))
	tagId := cast.ToInt(c.DefaultPostForm("tag_id", ""))
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	content := c.PostForm("content")
	modifiedBy := c.PostForm("modified_by")

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		article := models.GetArticle(id)
		data := models.CommonMaps()
		if article.ID > 0 {
			if models.ExistTagByID(tagId) {
				if tagId > 0 {
					data["tagId"] = tagId
				} else {
					data["tagId"] = article.TagID
				}
				if title != "" {
					data["title"] = title
				} else {
					data["title"] = article.Title
				}
				if desc != "" {
					data["Desc"] = desc
				} else {
					data["Desc"] = article.Desc
				}
				if content != "" {
					data["Content"] = content
				} else {
					data["Content"] = article.Content
				}
				if modifiedBy != "" {
					data["modified_by"] = modifiedBy
				} else {
					data["modified_by"] = article.ModifiedBy
				}
				fmt.Printf("%+v\n", data)

				models.EditArticle(id, data)
				code = e.SUCCESS
			} else {
				code = e.ERROR_NOT_EXIST_TAG
			}
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		fmt.Printf("%+v\n", valid.Errors)
		for _, err := range valid.Errors {
			logging.Error(err.Key + "=>" + err.Message)
		}
	}

	res := a.ToJSON(code, nil)
	c.JSON(http.StatusOK, res)
}

// @Summary Delete article
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/articles/{id} [delete]
func (a ArticleControoler) DeleteArticle(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistArticleByID(id) {
			models.DeleteArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logging.Error(err.Key + "=>" + err.Message)
		}
	}
	res := a.ToJSON(code, nil)
	c.JSON(http.StatusOK, res)
}
