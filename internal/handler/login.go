package handler

import (
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context){

	var (
		userInfo struct{
			Username string `json:"username"`
			Password string	`json:"password"`
		}
		rsp struct {
			Msg string `json:"msg"` 
			Token string `json:"token"`
		}
	)
	if err := c.ShouldBindJSON(&userInfo);err != nil {
		rsp.Msg = "登录参数错误..."
		return
	}

	rsp.Msg,rsp.Token = model.Login(userInfo.Username,userInfo.Password)

	defer func(){
		c.JSON(http.StatusOK,rsp)
	}()
}
