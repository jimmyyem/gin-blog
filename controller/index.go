package controller

import (
	"gin-blog/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

func Index(c *gin.Context) {
	offset := cast.ToInt(c.DefaultQuery("offset", "0"))
	limit := cast.ToInt(c.DefaultQuery("limit", "10"))

	// 获取列表和总数
	list := models.GetArticles(offset, limit, models.CommonMaps())
	total := models.GetArticleTotal(models.CommonMaps())

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "首页",
		"list":  list,
		"total": total,
	})
}
