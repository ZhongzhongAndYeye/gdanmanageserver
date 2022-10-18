package router

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

// token鉴权中间件
func JWTVerification(c *gin.Context) {
	var rsp struct {
		Msg        string `json:"msg"`
		Tokenvalid int    `json:"tokenvalid"` // token是否过期，0表示失效 1表示未失效
	}
	token := c.Request.Header.Get("token")
	_, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("iamlulu"), nil
	})
	if err != nil {
		rsp.Msg = "token已失效，请重新登录！"
		rsp.Tokenvalid = 0
		c.JSON(http.StatusOK, rsp)
		c.Abort()
	}
	c.Next()
}
