package handler

import (
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func Getxsskc(c *gin.Context){
	var rsp struct {
		Msg string
		Tokenvalid int
		Data []model.Xssdataall
	}
	rsp.Tokenvalid = 1
	rsp.Msg,rsp.Data = model.Getxsskc()
	defer func(){
		c.JSON(http.StatusOK,rsp)
	}()
}