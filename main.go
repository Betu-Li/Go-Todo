package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Todo Model
type todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func main() {
	// 创建数据库
	// sql:CREATE DATABASE Gin_Todo
	// 连接数据库

	router := gin.Default()
	// 设置静态文件位置
	router.Static("/static", "./static")
	// 设置模版文件位置
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// v1
	v1Group := router.Group("/v1")
	{
		// 添加
		v1Group.POST("/todo", func(c *gin.Context) {

		})
		// 查看所有待办事项
		v1Group.GET("/todo", func(c *gin.Context) {

		})
		// 查看某一条待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		// 修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {

		})
		// 删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {

		})
	}

	router.Run(":8080")
}
