package controller

import (
	"github.com/gin-gonic/gin"
	"go_code/ginStudy/gin01_b/bubble/dao"
	"go_code/ginStudy/gin01_b/bubble/models"
	"net/http"
)

/*
url --> controller --> logic --> model
请求 -->   控制器   -->业务逻辑 --> 模型层的增删改查
*/

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodoTodo(c *gin.Context) {
	// 前端页面填写待办事项 点击提交 发请求到这里
	// 1、从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	// 2、存入数据库
	err := dao.CreateOneTodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
		//c.JSON(http.StatusOK, gin.H{
		//	"code": 233333,
		//	"msg" : "success",
		//	"data": todo,
		//})
	}
}
func GetTodoList(c *gin.Context) {
	// 查询todo这个表里的所有数据
	todoList, err := dao.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}
func UpdateOneTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	todo, err := dao.GetOneTodo(id)
	c.BindJSON(&todo)
	if err = dao.UpdateOneTodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteOneTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := dao.DeleteOneTodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
