package controller

import (
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexController(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateATodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	if err := models.CreateATodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(c *gin.Context) {
	// 查询todos表里的所有数据
	var todoList []models.Todo
	if err := models.GetTodoList(&todoList); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "id 不存在",
		})
		return
	}
	var todo models.Todo
	if err := models.UpdateATodo(id, &todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&todo)
}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "id 不存在",
		})
		return
	}
	if err := models.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
}
