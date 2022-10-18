package model

import (
	"server/utils"
)

type Table struct {
	Id        string `json:"id"`
	Tablename string `json:"tablename"`
}

func Getxsstable() (status int, tablelist []Table) {
	db := utils.DB
	err := db.Table("xsstablelist").Select("id,tablename").Order("put_in_time desc").Find(&tablelist).Error
	if err != nil {
		status = 0
		return
	}
	status = 1
	return
}
