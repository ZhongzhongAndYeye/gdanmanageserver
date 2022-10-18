package handler

import (
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func Getxsstable(c *gin.Context) {
	var rsp struct {
		Status       int           `json:"status"`
		Tokenvalid   int           `json:"tokenvalid"`
		Xsstablelist []model.Table `json:"xsstablelist"`
	}
	rsp.Status, rsp.Xsstablelist = model.Getxsstable()
	rsp.Tokenvalid = 1
	defer func() {
		c.JSON(http.StatusOK, rsp)
	}()
}
