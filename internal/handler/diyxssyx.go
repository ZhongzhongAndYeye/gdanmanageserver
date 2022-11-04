package handler

import (
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func Diyxssyx(c *gin.Context) {
	var req struct {
		Tableid   string
		Tablepjid string
		Data      []model.Pjdataslice
	}
	var rsp struct {
		Tokenvalid int `json:"tokenvalid"`
		Status     int `json:"status"` 		// 1表示成功 2表示此桌子上存在相同牌局 3表示牌库中存在相同牌局 0表示出现错误
	}
	rsp.Tokenvalid = 1
	c.ShouldBindJSON(&req)
	rsp.Status = model.Diyxssyx(req.Tableid,req.Tablepjid,req.Data)
	defer func(){
		c.JSON(http.StatusOK,rsp)
	}()
}
