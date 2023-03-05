package main

import (
	"github.com/gin-gonic/gin"
	viper2 "github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"todo/dao"
	"todo/middleware"
)

func main() {
	v := viper2.New()
	v.SetConfigFile("secret.yml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	dsn := v.GetString("Dsn")

	//初始化数据库
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	dao.SetDefault(db)

	//初始化路由
	g := gin.New()
	g.Use(gin.Logger())
	g.Use(gin.Recovery())
	g.Use(middleware.Cors())
	g.Use(middleware.Version())
	InitRouter(g)

	//err2 := g.Run(":8081")
	err = g.RunTLS(":8081", "cloud.jmzzz.cn.crt", "cloud.jmzzz.cn.key")
	if err != nil {
		return
	}
}
