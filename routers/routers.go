package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 静态文件路径
	r.Static("/static", "static")
	// 模板文件路径
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.IndexController)

	v1Group := r.Group("/v1")
	{
		// 增加代办事项
		v1Group.POST("/todo", controller.CreateATodo)

		// 查询所有代办事项
		v1Group.GET("/todo", controller.GetTodoList)

		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)

		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
