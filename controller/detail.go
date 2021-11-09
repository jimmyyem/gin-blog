package controller

import (
	"fmt"
	"gin-blog/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
)

func Detail(c *gin.Context) {
	id := cast.ToInt(c.Query("id"))
	article := models.Article{}

	if id > 0 {
		article = models.GetArticle(id)
	}
	fmt.Printf("%+v", article)

	c.HTML(http.StatusOK, "detail.html", gin.H{
		"title":   "文章详情",
		"article": article,
	})
}
