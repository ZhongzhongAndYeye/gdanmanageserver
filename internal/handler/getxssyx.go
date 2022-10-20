package handler

import (
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func Getxssyx(c *gin.Context) {
	var (
		tableid struct {
			Tableid string `form:"tableid"`
		}
		rsp struct {
			Status     int               `json:"status"` // 1表示成功 2表示获取数据为空 0表示请求失败
			Tokenvalid int               `json:"tokenvalid"`
			Data       []model.Xssyxdata `json:"data"`
		}
	)
	c.ShouldBind(&tableid)
	rsp.Status, rsp.Data = model.Getxssyx(tableid.Tableid)
	rsp.Tokenvalid = 1
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()
}
