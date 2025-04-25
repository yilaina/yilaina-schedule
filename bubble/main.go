package main

import (
	"go_code/ginStudy/gin01_b/bubble/dao"
	"go_code/ginStudy/gin01_b/bubble/routers"
)

func main() {
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.CloseMySQL() // 程序退出关闭数据库连接
	// 模型绑定
	dao.InitModel()
	// 注册路由
	r := routers.SetupRouter()
	r.Run(":8080")
}
