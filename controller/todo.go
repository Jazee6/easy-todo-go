package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo/model"
	"todo/service"
)

func GetTodoList(c *gin.Context) {
	v, _ := c.Get("userid")
	list, err := service.GetTodoList(v.(float64))
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code: -1,
			Msg:  "获取todo列表失败",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: 0,
		Msg:  "",
		Data: list,
	})
}

type createTodoReq struct {
	Todo string `json:"todo" binding:"required"`
}

func CreateTodo(c *gin.Context) {
	var req createTodoReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: -1,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}

	v, _ := c.Get("userid")
	if err := service.AddTodoItem(int32(v.(float64)), req.Todo); err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code: -1,
			Msg:  "添加失败",
			Data: nil,
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Code: 0,
		Msg:  "添加成功",
		Data: nil,
	})
}
