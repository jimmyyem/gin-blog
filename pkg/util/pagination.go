package util

import (
	"gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func GetOffset(c *gin.Context) int {
	result := 0
	page := cast.ToInt(c.Query("page"))
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}
