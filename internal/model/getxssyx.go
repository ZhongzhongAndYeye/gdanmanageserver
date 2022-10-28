package model

import (
	"fmt"
	"server/utils"
)

type Xssyxdata struct {
	Tablepjid   string   `json:"tablepjid"`
	Tablepjdata []Pjdata `json:"tablepjdata"`
}

func Getxssyx(tableid string) (status int, xssyxdata []Xssyxdata) {
	db := utils.DB
	// 获取桌子的桌子牌局id列表
	type Tablepj struct {
		Tablepjid string
	}
	tablepjs := []Tablepj{}
	err := db.Table("xsstablepj").Where("tableid = ?", tableid).Find(&tablepjs).Error
	if err != nil {
		fmt.Println("操作xsstablepj时出现错误...", err)
		status = 0
		return
	}
	// 根据桌子上的牌局数量来开辟获取数组空间
	num := len(tablepjs)
	xssyxdata = make([]Xssyxdata, num)
	for i, v := range tablepjs {
		xssyxdata[i].Tablepjid = v.Tablepjid
		err = db.Table("xsstablepjdatas").Where("tablepjid = ?", v.Tablepjid).Order("`order` asc").Find(&xssyxdata[i].Tablepjdata).Error
		if err != nil {
			fmt.Println("操作xsstablepjdatas表时发生错误", err)
			status = 0
			return
		}
	}
	if len(xssyxdata) == 0 {
		status = 2
		return
	}
	status = 1
	return
}
