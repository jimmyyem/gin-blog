package controller

import (
	"gin-blog/pkg/e"
	"github.com/gin-gonic/gin"
)

type BaseController struct {
	ctx gin.Context
}

// 构造返回内容方法
func (b BaseController) ToJSON(code e.ErrCode, data interface{}) map[string]interface{} {
	return map[string]interface {
	}{
		"code": int(code),
		"msg":  code.String(),
		"data": data,
	}
}
