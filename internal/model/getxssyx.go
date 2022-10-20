package model

import (
	"fmt"
	"server/utils"
)

type Xssyxdata struct {
	Pjid   string   `json:"pjid"`
	Pjdata []Pjdata `json:"pjdata"`
}

func Getxssyx(tableid string) (status int, xssyxdata []Xssyxdata) {
	db := utils.DB
	type Tablepj struct {
		Pjid string
	}
	tablepjs := []Tablepj{}
	err := db.Table("xsstablepj").Where("tableid = ?", tableid).Find(&tablepjs).Error
	if err != nil {
		fmt.Println("操作xsstablepj时出现错误...", err)
		status = 0
		return
	}
	num := len(tablepjs)
	xssyxdata = make([]Xssyxdata, num)
	for i, v := range tablepjs {
		xssyxdata[i].Pjid = v.Pjid
		err = db.Table("xsstablepjdatas").Where("pjid = ?", v.Pjid).Find(&xssyxdata[i].Pjdata).Error
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
