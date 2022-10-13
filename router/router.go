package router

import (
	"server/internal/handler"
	"github.com/gin-gonic/gin"
)

func RouterStart() {
	r := gin.Default()

	r.POST("/sessions", handler.Login) // 登录
	// 后台管理接口
	g := r.Group("/guandanmanage")
	g.Use(JWTVerification)
	g.GET("/xsskc",handler.Getxsskc) // 查询限时赛所有牌局库存数据
	
	// 客户端接口
	c := r.Group("/client")
	c.POST("/xsspaijudata",handler.Xsspaijuinsert) // 客户端牌局数据入库

	r.Run(":1234")
}
