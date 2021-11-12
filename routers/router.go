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

	// 设置静态页面目录
	r.Static("/static", "./static")

	// 设置swagger路径
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 模板信息
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
	auth := api.AuthController{}
	pc.Any("/auth", auth.GetAuth)

	index := controller.IndexController{}
	pc.GET("/", index.Index)
	pc.GET("/index", index.Index)
	pc.GET("/detail", index.Detail)
}

func addGroupRouter(apiv1 *gin.RouterGroup) {
	apiv1.Use(jwt.JWT())
	{
		tag := v1.TagController{}
		//获取标签列表
		apiv1.GET("/tags", tag.GetTags)
		//获取标签详情
		apiv1.GET("/tags/:id", tag.GetTag)
		//新建标签
		apiv1.POST("/tags", tag.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", tag.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", tag.DeleteTag)

		art := v1.ArticleControoler{}
		//获取文章列表
		apiv1.GET("/articles", art.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", art.GetArticle)
		//新建文章
		apiv1.POST("/articles", art.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", art.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", art.DeleteArticle)
	}
}
