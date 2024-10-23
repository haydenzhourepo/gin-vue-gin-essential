package main

import (
	"github.com/gin-gonic/gin"
	"oceanlearn.teach/ginessential/controller"
	"oceanlearn.teach/ginessential/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	// 创建路由的api分组
	api := r.Group("/api/")
	{
		// 在api分组里面创建auth分组
		auth := api.Group("/auth")
		{
			// 注册接口
			auth.POST("/register", controller.Register)
			// 登录接口
			auth.POST("/login", controller.Login)
			// 获取用户信息，此接口需要鉴权中间件认证
			auth.GET("/info", middleware.AuthMiddleware(), controller.Info)
		}
	}

	// 创建categoryRoutes分组
	categoryRoutes := r.Group("/categories")
	{
		// 创建categoryController
		categoryController := controller.NewCategoryController()
		// category创建接口
		categoryRoutes.POST("", categoryController.Create)
		// category更新接口
		categoryRoutes.PUT(":id", categoryController.Update)
		// category展示接口
		categoryRoutes.GET(":id", categoryController.Show)
		// category删除接口
		categoryRoutes.DELETE(":id", categoryController.Delete)
	}

	// 创建postRoutes分组
	postRoutes := r.Group("/posts")
	{
		// postRoute使用鉴权中间件
		postRoutes.Use(middleware.AuthMiddleware())
		// 创建postController
		postController := controller.NewPostController()
		// post创建接口
		postRoutes.POST("", postController.Create)
		// post更新接口
		postRoutes.PUT(":id", postController.Update)
		// post展示接口
		postRoutes.GET(":id", postController.Show)
		// post删除接口
		postRoutes.DELETE(":id", postController.Delete)
		// 分页接口
		postRoutes.POST("page/list", postController.PageList)
	}
	return r
}
