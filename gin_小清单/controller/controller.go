package controller

import (
	"bubble/dao"
	"bubble/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	//渲染模板文件
	c.HTML(http.StatusOK, "index.html", nil)

}
func AndDB(c *gin.Context) {
	//从请求中获取数据
	var todo models.Todo
	c.BindJSON(&todo)
	//存入数据库
	dao.DB.Create(&todo)
	//返回相应
	//err...........

}
func LOOK(c *gin.Context) {
	//查询Todo这个表里的所有数据
	var todolist []models.Todo
	dao.DB.Find(&todolist)
	c.JSON(http.StatusOK, todolist)
}
func Change(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var todo models.Todo
	dao.DB.Where("id=?", id).First(&todo)
	c.BindJSON(&todo)
	dao.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)

}
func DeletcDB(c *gin.Context) {
	id, _ := c.Params.Get("id")
	dao.DB.Where("id=?", id).Delete(models.Todo{})
}
