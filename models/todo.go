package models

import (
	"bubble/dao"
)

type Todo struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// 一般放todo类, 并且其增删改查操作

func CreateATodo(todo *Todo) (err error) {
	if err = dao.DB.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

func GetTodoList(todoList *[]Todo) (err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return err
	}
	return nil
}

func UpdateATodo(id string, todo *Todo) (err error) {
	if err = dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return err
	}
	if err = dao.DB.Save(&todo).Error; err != nil {
		return err
	}
	return nil
}

func DeleteATodo(id string) (err error) {
	if err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error; err != nil {
		return err
	}
	return nil
}
