package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin-gorm/controller"
	"go-gin-gorm/middleware"
)

// CollectRoutes 路由
func CollectRoutes(r *gin.Engine) *gin.Engine {
	// 允许跨域访问
	r.Use(middleware.CORSMiddleware())
	// 注册
	r.POST("/register", controller.Register)
	// 登录
	r.POST("/login", controller.Login)
	// 登录获取用户信息
	r.GET("/user", middleware.AuthMiddleware(), controller.GetInfo)
	// 上传图像
	r.POST("/upload", controller.Upload)
	// 查询分类
	r.GET("/category", controller.SearchCategory)         // 查询分类
	r.GET("/category/:id", controller.SearchCategoryName) // 查询分类名
	//用户文章的增删查改
	articleRoutes := r.Group("/article")
	articleController := controller.NewArticleController()
	articleRoutes.POST("", middleware.AuthMiddleware(), articleController.Create)      // 发布文章
	articleRoutes.PUT(":id", middleware.AuthMiddleware(), articleController.Update)    // 修改文章
	articleRoutes.DELETE(":id", middleware.AuthMiddleware(), articleController.Delete) // 删除文章
	articleRoutes.GET(":id", articleController.Show)                                   // 查看文章
	articleRoutes.POST("list", articleController.List)

	return r
}
