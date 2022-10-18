package model

import "server/utils"

func Delxsskc(pjid string) (msg string) {
	type Xsspaijudata struct {
		Id        int
		Pjid      string
		Order     int
		HandCard  string
	}
	type Xsspjid struct {
		Pjid string
		PutInTime string
		Frequency float32
		Count     int
	}
	db := utils.DB
	err := db.Table("xsspjidlist").Where("pjid = ?",pjid).Delete(Xsspjid{}).Error
	if err != nil{
		msg = "操作xsspjidlist表失败，删除记录失败！"
		return
	}
	err = db.Table("xsspaijudatas").Where("pjid = ?",pjid).Delete(Xsspaijudata{}).Error
	if err != nil{
		msg = "操作xsspaijudatas表失败，删除记录失败！"
		return
	}
	msg = "删除记录成功"
	return
}