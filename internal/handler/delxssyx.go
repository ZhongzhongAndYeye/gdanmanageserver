package handler

import (
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func Delxssyx(c *gin.Context){
	var tablepj struct {
		Tableid string
		Pjid string
	}
	var rsp struct {
		Status int `json:"status"`
		Tokenvalid int `json:"tokenvalid"`
	}
	rsp.Tokenvalid = 1
	c.ShouldBindJSON(&tablepj)
	rsp.Status = model.Delxssyx(tablepj.Tableid,tablepj.Pjid)
	defer func(){
		c.JSON(http.StatusOK,rsp)
	}()
}