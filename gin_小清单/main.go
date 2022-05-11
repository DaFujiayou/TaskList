package main

import (
	"bubble/controller"
	"bubble/dao"
	"bubble/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	//创建数据库
	//连接数据库
	dao.Initmysql()
	//程序退出前 关闭数据库
	defer dao.DB.Close()
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	r := gin.Default()
	//寻找模板文件中引用的静态文件（js,css等文件）
	r.Static("./static", "static")
	//解析模板文件
	r.LoadHTMLGlob("template/*")

	r.GET("/", controller.IndexHandler) //此处的controller

	//定义API
	//v1
	v1Group := r.Group("v1")
	//待办事项模块
	//增
	v1Group.POST("/todo", controller.AndDB)
	//查
	//查看所有待办事项
	v1Group.GET("/todo", controller.LOOK)
	//改
	v1Group.PUT("/todo/:id", controller.Change)

	//删
	v1Group.DELETE("/todo/:id", controller.DeletcDB)

	r.Run(":8088")

}
