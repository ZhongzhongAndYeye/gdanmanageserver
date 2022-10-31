package handler

import (
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func Changexssyxorder(c *gin.Context) {
	var req struct {
		Tablepjid string
		Beforeo   int
		Aftero    int
		Tableid   string
	}
	var rsp struct {
		Tokenvalid int `json:"tokenvalid"`
		Status     int `json:"status"`
	}
	c.ShouldBindJSON(&req)
	rsp.Tokenvalid = 1
	rsp.Status = model.Changexssyxorder(req.Tablepjid, req.Beforeo, req.Aftero,req.Tableid)
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()
}
