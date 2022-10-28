package handler

import (
	"fmt"
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func Delxssyx(c *gin.Context) {
	var tablepj struct {
		Tablepjid string
	}
	var rsp struct {
		Status     int `json:"status"`
		Tokenvalid int `json:"tokenvalid"`
	}
	rsp.Tokenvalid = 1
	c.ShouldBindJSON(&tablepj)
	fmt.Println(tablepj)
	rsp.Status = model.Delxssyx(tablepj.Tablepjid)
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()
}
