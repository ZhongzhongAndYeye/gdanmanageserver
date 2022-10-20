package handler

import (
	"fmt"
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

// 限时赛出库
func Outofxsskc(c *gin.Context) {
	var (
		xssck struct {
			Tableid string
			Pjid    string
		}
		rsp struct {
			Status     int `json:"status"`
			Tokenvalid int `json:"tokenvalid"`
		}
	)
	c.ShouldBindJSON(&xssck)
	fmt.Println(xssck)
	rsp.Status = model.Outofxsskc(xssck.Tableid, xssck.Pjid)
	rsp.Tokenvalid = 1
	defer func(){
		c.JSON(http.StatusOK,rsp)
	}()
}
