package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo/model"
	"todo/service"
)

type loginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResp struct {
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: -1,
			Msg:  "登录参数错误",
			Data: nil,
		})
		return
	}

	token, err := service.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code: -1,
			Msg:  "用户名或密码错误",
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: 0,
		Msg:  "登录成功",
		Data: loginResp{Token: token},
	})
}

type registerReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type regResp struct {
	Token string `json:"token"`
}

func Register(c *gin.Context) {
	var req registerReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: -1,
			Msg:  "注册参数错误",
			Data: nil,
		})
		return
	}

	token, err := service.Register(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			Code: -1,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: 0,
		Msg:  "注册成功",
		Data: regResp{Token: token},
	})
}
