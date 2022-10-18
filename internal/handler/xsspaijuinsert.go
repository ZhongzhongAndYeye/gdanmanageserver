package handler

import (
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func Xsspaijuinsert(c *gin.Context) {
	var (
		paijudata struct {
			Data [][]int `json:"data"`
		}

		rsp struct{
			Msg string	`json:"msg"`
		}
	)
	c.ShouldBindJSON(&paijudata)
	rsp.Msg = model.Xsspaijuinsert(paijudata.Data)
	defer func() {
		c.JSON(http.StatusOK,rsp)
	}()
}
