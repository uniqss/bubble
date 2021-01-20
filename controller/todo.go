package controller

import (
	"bubble/models"
	Util "bubble/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func TodoCreate(c *gin.Context) {
	log := Util.ZLog
	log.Debug("TodoCreate", zap.String("goid", Util.GetCurrentGoroutineIdStr()))
	// 前端页面填写待办事项 点击提交 会发请求到这里
	// 1. 从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	// 2. 存入数据库
	err := models.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
		//c.JSON(http.StatusOK, gin.H{
		//	"code": 2000,
		//	"msg": "success",
		//	"data": todo,
		//})
	}
}

func TodoGetAll(c *gin.Context) {
	log := Util.ZLog
	log.Debug("TodoGetAll", zap.String("goid", Util.GetCurrentGoroutineIdStr()))
	// 查询todo这个表里的所有数据
	todoList, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func TodoUpdate(c *gin.Context) {
	log := Util.ZLog
	log.Debug("TodoUpdate", zap.String("goid", Util.GetCurrentGoroutineIdStr()))
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err = models.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func TodoDelete(c *gin.Context) {
	log := Util.ZLog
	log.Debug("TodoDelete", zap.String("goid", Util.GetCurrentGoroutineIdStr()))
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := models.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
