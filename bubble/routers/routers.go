package routers

import (
	"github.com/gin-gonic/gin"
	"go_code/ginStudy/gin01_b/bubble/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.IndexHandler)
	// v1
	v1Group := r.Group("/v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodoTodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateOneTodo)
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteOneTodo)
	}
	return r
}
