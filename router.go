package main

import (
	"github.com/gin-gonic/gin"
	"todo/controller"
	"todo/middleware"
)

func InitRouter(r *gin.Engine) {
	api := r.Group("/todo")

	api.POST("/login", controller.Login)
	api.POST("/register", controller.Register)

	api.Use(middleware.Auth())
	{
		api.GET("/todos", controller.GetTodoList)
		api.POST("/todos", controller.CreateTodo)
	}

}
