package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

// Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// 连接数据库
func initMySql() (err error) {
	dsn := "root:mysql123@tcp(127.0.0.1:3306)/gin_todo?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = DB.DB().Ping()
	return
}

func main() {
	// 创建数据库
	// sql:CREATE DATABASE gin_Todo
	// 连接数据库
	err := initMySql()
	if err != nil {
		panic(err)
	}
	defer DB.Close() //程序退出时关闭数据库
	// 绑定数据库
	DB.AutoMigrate(&Todo{})

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
			// 填写代办事项，点击提交
			var todo Todo
			// JSON绑定
			if err := c.ShouldBindJSON(&todo); err != nil {
				c.JSON(http.StatusOK, gin.H{"err": "Invalid input data"})
				return
			}
			// 在数据库中创建数据
			err := DB.Create(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 2000,
					"msg":  "success",
					"data": todo,
				})
			}
		})
		// 查看所有待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []Todo
			err := DB.Find(&todoList).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todoList)
			}

		})
		// 查看某一条待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		// 修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			// 获取记录id信息
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"err": "无效id"})
				return
			}
			var todo Todo
			// 查询数据库中是否有这条记录
			err := DB.Where("id = ?", id).First(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"err": err.Error()})
				return
			}
			// 从前端获取修改后的status值，JSON绑定
			var updatedData struct {
				Status string `json:"status"`
			}
			if err := c.ShouldBindJSON(&updatedData); err != nil {
				println("JSON绑定错误:", err.Error())
				c.JSON(http.StatusOK, gin.H{"err": "Invalid input data"})
				return
			}
			// 打印status到控制台
			println(todo.Status)
			if err := DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"err": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		// 删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			// 拿到记录id信息
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"err": "无效id"})
				return
			}

			// 通过id找到记录并删除
			err := DB.Where("id=?", id).Delete(Todo{}).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"err": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"id":   "delete",
				"code": 2000,
				"msg":  "success",
			})
		})
	}

	router.Run(":8080")
}
