package models

import "Gin-Todo/setting"

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// CreateTodo 新增代办事项
func CreateTodo(todo *Todo) (err error) {
	err = setting.DB.Create(todo).Error
	return
}

// GetTodos 获取所有代办事项】
func GetTodos() (todoList []*Todo, err error) {
	err = setting.DB.Find(&todoList).Error
	if err != nil {
		return nil, err
	}
	return
}

// GetTodo 获取单个代办事项
func GetTodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	err = setting.DB.Debug().Where("id=?", id).First(todo).Error
	if err != nil {
		return nil, err
	}
	return
}

// UpdateTodo 更新代办事项
func UpdateTodo(todo *Todo) (err error) {
	err = setting.DB.Save(todo).Error
	return
}

// DeleteTodo 删除代办事项
func DeleteTodo(id string) (err error) {
	err = setting.DB.Where("id=?", id).Delete(&Todo{}).Error
	if err != nil {
		return err
	}
	return
}
