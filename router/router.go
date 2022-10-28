package router

import (
	"github.com/gin-gonic/gin"
	"server/internal/handler"
)

func RouterStart() {
	r := gin.Default()

	r.POST("/sessions", handler.Login) // 登录
	// 后台管理接口
	g := r.Group("/gdanmanage")
	g.Use(JWTVerification)
	g.GET("/xsskc", handler.Getxsskc)                // 查询限时赛所有牌局库存数据
	g.DELETE("/xsskc", handler.Delxsskc)             // 删除指定的限时赛库存牌局
	g.POST("/xsskc", handler.Outofxsskc)             // 限时赛牌局出库
	g.POST("/xsstable", handler.Createxsstable)      // 新建限时赛桌子
	g.GET("/xsstable", handler.Getxsstable)          // 获取限时赛桌子列表
	g.GET("/xssyx", handler.Getxssyx)                // 查询限时赛指定桌子上的牌局信息
	g.DELETE("/xssyx", handler.Delxssyx)             // 删除指定桌子上的指定牌局
	g.POST("/xssyx/order", handler.Changexssyxorder) // 改变指定牌局的order属性

	// 客户端接口
	c := r.Group("/client")
	c.POST("/xsspaijudata", handler.Xsspaijuinsert) // 客户端牌局数据入库

	r.Run(":1234")
}
