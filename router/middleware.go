package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

// token鉴权中间件
func JWTVerification(c *gin.Context) {
	var tokenreq struct{
		Token string
	}
	var rsp struct{
		Msg string
		Tokenvalid int // token是否过期，0表示过期 1表示未过期
	}
	c.ShouldBindJSON(&tokenreq)
	fmt.Println(tokenreq)
	token, err := jwt.ParseWithClaims(tokenreq.Token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("iamlulu"), nil
	})
	if err != nil {
		fmt.Println("err:",err)
	}
	fmt.Println(token.Claims.(*jwt.StandardClaims)) // 类型断言成我们所需要的，注意断言成指针
	t := token.Claims.(*jwt.StandardClaims) 
	fmt.Println(t.ExpiresAt)
	now := time.Now().Unix()
	if now <= t.ExpiresAt{
		c.Next()
	}else{
		rsp.Msg = "登录已过期，请重新登陆！"
		rsp.Tokenvalid = 0
		c.JSON(http.StatusOK,rsp)
	}

}