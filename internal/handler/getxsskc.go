package handler

import (
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func Getxsskc(c *gin.Context) {
	var rsp struct {
		Msg        string             `json:"msg"`
		Tokenvalid int                `json:"tokenvalid"`
		Data       []model.Xssdataall `json:"data"`
	}
	rsp.Tokenvalid = 1
	rsp.Msg, rsp.Data = model.Getxsskc()
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()
}
