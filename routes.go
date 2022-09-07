package main


//创建路由

import (
	"github.com/gin-gonic/gin"
	"xukun/MiddleWare"
	"xukun/cotroller"
)

func ControllerRouter(r *gin.Engine) *gin.Engine {
	r.Use(MiddleWare.CORSMiddleware())
	r.POST("/api/auth/register", cotroller.Register) //func(c *gin.Context) 是为r.GET这个方法提供具体的操作
	r.POST("/api/auth/login", cotroller.Login)
	//用中间件保护用户信息接口
	r.GET("/api/auth/info", MiddleWare.AuthMiddleWare(), cotroller.Info)

	//用一个路由分组实现文章的CRUD
	CategoryRoutes := r.Group("/categories")
	//对CategoryRoutes实例化
	CategoryController := cotroller.NewCategoryController()
	//增
	CategoryRoutes.POST("", CategoryController.Create)
	//改
	CategoryRoutes.PUT("/:id", CategoryController.Update)
	//查
	CategoryRoutes.GET("/:id", CategoryController.Show)
	//删
	CategoryRoutes.DELETE("/:id", CategoryController.Delete)
	//PATCH对部分内容进行更改
	return r
}