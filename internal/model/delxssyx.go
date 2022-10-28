package model

import (
	"fmt"
	"server/utils"
)

func Delxssyx(tablepjid string) (status int) {
	db := utils.DB
	type Xsstablepj struct {
		Id        int
		Tableid   string
		Tablepjid string
	}
	err := db.Table("xsstablepj").Where("tablepjid = ?", tablepjid).Delete(Xsstablepj{}).Error
	if err != nil {
		fmt.Println("操作xsstablepj时出现错误...", err)
		status = 0
		return
	}
	type Xsstablepjdatas struct {
		Id        int
		Tablepjid string
		Order     int
		HandCard  string
	}
	err = db.Table("xsstablepjdatas").Where("tablepjid = ?", tablepjid).Delete(Xsstablepjdatas{}).Error
	if err != nil {
		fmt.Println("操作xsstablepjdatas时出现错误...", err)
		status = 0
		return
	}
	status = 1
	return
}
