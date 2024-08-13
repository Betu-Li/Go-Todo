package v1

import (
	"Gin-Todo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// CreateTodo 添加代办事项
func CreateTodo(c *gin.Context) {
	var todo models.Todo
	// JSON绑定
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"err": "输入数据无效",
		})
		return
	}
	if err := models.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": todo,
		})
	}
}

// GetTodos 查看所有代办事项
func GetTodos(c *gin.Context) {
	if todoList, err := models.GetTodos(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})

	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

// UpdateTodo 更新待办事项状态
func UpdateTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"err": "无效id"})
		return
	}
	// 查询数据库中是否有这一条记录
	todo, err := models.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	}
	if err = c.BindJSON(&todo); err != nil {
		println("JSON绑定错误:", err.Error())
		c.JSON(http.StatusOK, gin.H{"err": "输入数据无效"})
		return
	}
	// 在数据库中更新
	if err = models.UpdateTodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"err": "无效id"})
		return
	}
	// 查询数据库中是否有这一条记录
	_, err := models.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	}
	if err = models.DeleteTodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":   "delete",
		"code": 2000,
		"msg":  "success",
	})
}
