package model

import (
	"fmt"
	"server/utils"
)

func Changexssyxorder(tablepjid string, beforeo int, aftero int) (status int) {
	db := utils.DB
	err := db.Table("xsstablepjdatas").Where("tablepjid = ?",tablepjid).Where("`order` = ?",aftero).Update("order",666).Error
	if err != nil {
		fmt.Println("操作xsstablepjdatas时出现错误1...",err)
		status = 0
		return
	}
	err = db.Table("xsstablepjdatas").Where("tablepjid = ?",tablepjid).Where("`order` = ?",beforeo).Update("order",aftero).Error
	if err != nil {
		fmt.Println("操作xsstablepjdatas时出现错误2...",err)
		status = 0
		return
	}
	err = db.Table("xsstablepjdatas").Where("tablepjid = ?",tablepjid).Where("`order` = ?",666).Update("order",beforeo).Error
	if err != nil {
		fmt.Println("操作xsstablepjdatas时出现错误3...",err)
		status = 0
		return
	}
	status = 1
	return
}