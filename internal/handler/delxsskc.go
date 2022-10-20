package handler

import (
	"fmt"
	"net/http"
	"server/internal/model"

	"github.com/gin-gonic/gin"
)

func Delxsskc(c *gin.Context) {
	var (
		pjidcon struct {
			Pjid string
		}
		rsp struct {
			Msg        string `json:"msg"`
			Tokenvalid int    `json:"tokenvalid"`
		}
	)
	c.ShouldBindJSON(&pjidcon)
	rsp.Tokenvalid = 1
	fmt.Println(pjidcon)
	rsp.Msg = model.Delxsskc(pjidcon.Pjid)
	defer func ()  {
		c.JSON(http.StatusOK,rsp)
	}()
}
