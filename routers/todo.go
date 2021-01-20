package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouterTodo(r *gin.Engine) {
	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.TodoCreate)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.TodoGetAll)
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.TodoUpdate)
		// 删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.TodoDelete)
	}
}
