package routers

import (
	"gin-blog/controller"
	"gin-blog/controller/api"
	"gin-blog/controller/api/v1"
	_ "gin-blog/docs"
	"gin-blog/middleware/jwt"
	"gin-blog/pkg/setting"
	"gin-blog/pkg/util"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"html/template"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Delims("{{", "}}")
	r.SetFuncMap(template.FuncMap{
		"formatAsTime": util.FormatAsTime,
	})
	r.LoadHTMLGlob("template/*.html")

	// pc网页路由
	pcGroup := r.Group("")
	addGroupPc(pcGroup)

	// api路由
	apiv1 := r.Group("/api/v1")
	addGroupRouter(apiv1)

	return r
}

func addGroupPc(pc *gin.RouterGroup) {
	pc.Any("/auth", api.GetAuth)
	pc.GET("/index", controller.Index)
	pc.GET("/detail", controller.Detail)
}

func addGroupRouter(apiv1 *gin.RouterGroup) {
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//获取标签详情
		apiv1.GET("/tags/:id", v1.GetTag)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}
}
