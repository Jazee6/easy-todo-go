package middleware

import (
	"github.com/gin-gonic/gin"
	"todo/config"
)

func Version() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Version", config.Cfg{}.GetVersion())
		c.Next()
	}
}
