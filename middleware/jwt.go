package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"time"
	"todo/model"
)

var secret = []byte("a9effa4d1b4046768d1dd31a52d4af33")

const exp = time.Hour * 24 * 30

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.Redirect(302, "/login")
			c.Abort()
			return
		}
		id := tokenValid(token)
		if id == -1 {
			c.Redirect(302, "/login")
			c.Abort()
			return
		}
		c.Set("userid", id)
		c.Next()
	}
}

func tokenValid(token string) float64 {
	tk, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return -1
	}
	if tk.Valid {
		parse, err := time.Parse(time.RFC3339, tk.Claims.(jwt.MapClaims)["myExp"].(string))
		if err != nil {
			return -1
		}
		if time.Now().After(parse) {
			_ = fmt.Errorf("token expired")
			return -1
		}
	}
	return tk.Claims.(jwt.MapClaims)["user"].(float64)
}

func GenToken(u model.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":  u.ID,
		"myExp": time.Now().Add(exp).Format(time.RFC3339),
	})
	signedString, err := token.SignedString(secret)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return signedString
}
