package service

import (
	"fmt"
	"regexp"
	"strings"
	"todo/dao"
	"todo/middleware"
	"todo/model"
)

func Login(username, password string) (string, error) {
	q := dao.Q.User
	u, err := q.Where(q.Username.Eq(username)).Where(q.Password.Eq(password)).Take()
	if err != nil {
		return "", err
	}

	return middleware.GenToken(*u), nil
}

func Register(username, password string) (string, error) {
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)
	if username == "" || password == "" {
		return "", fmt.Errorf("用户名或密码不能为空")
	}
	reg := regexp.MustCompile(`^[0-9a-zA-Z]\w{5,17}$`)
	if !reg.MatchString(password) {
		return "", fmt.Errorf("密码格式错误")
	}

	q := dao.Q.User
	if u, _ := q.Where(q.Username.Eq(username)).Take(); u != nil {
		return "", fmt.Errorf("用户已存在")
	}

	user := model.User{
		Username: username,
		Password: password,
	}
	if err := q.Create(&user); err != nil {
		return "", err
	}

	return middleware.GenToken(user), nil
}
