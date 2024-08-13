package routers

import (
	v1 "Gin-Todo/routers/controller/v1"
	"Gin-Todo/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	// 设置静态文件位置
	router.Static("/static", "static")
	// 设置模版文件位置
	router.LoadHTMLGlob("templates/*")

	// 设置路由
	router.GET("/", v1.IndexHandler)

	// v1
	v1Group := router.Group("/v1")
	{
		// 添加代办事项
		v1Group.POST("/todo", v1.CreateTodo)
		// 查看所有待办事项
		v1Group.GET("/todo", v1.GetTodos)
		// 更新待办事项状态
		v1Group.PUT("/todo/:id", v1.UpdateTodo)
		// 删除待办事项
		v1Group.DELETE("/todo/:id", v1.DeleteTodo)
	}

	return router
}
