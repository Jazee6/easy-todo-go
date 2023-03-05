package middleware

import (
	"github.com/gin-gonic/gin"
	"todo/config"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		list := config.Cfg{}.GetWhiteList()
		for _, v := range list {
			if c.Request.Header.Get("Origin") == v {
				c.Header("Access-Control-Allow-Origin", v)
				break
			}
		}
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Expose-Headers", "Version")

		if method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
