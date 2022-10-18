package handler

import (
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func Createxsstable(c *gin.Context) {
	var (
		tablename struct {
			Tablename string
		}
		rsp struct {
			Status     int `json:"status"`     // 创建是否成功 1表示成功 2表示桌子名重复操作失败 0表示其他失败
			Tokenvalid int `json:"tokenvalid"` // token是否过期 1表示有效 0表示过期
		}
	)
	c.ShouldBindJSON(&tablename)
	rsp.Status = model.Createxsstable(tablename.Tablename)
	rsp.Tokenvalid = 1
	defer func(){
		c.JSON(http.StatusOK,rsp)
	}()
}
