package model

import (
	"errors"
	"fmt"
	"server/utils"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func Createxsstable(tablename string) (status int) {

	type Table struct {
		Id        string
		Tablename string
		PutInTime int64
	}

	db := utils.DB
	tablenamesql := Table{}
	
	err := db.Table("xsstablelist").Where("tablename = ?", tablename).First(&tablenamesql).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 若错误是没有查询到同名数据，进行创建
			tablenamesql.Id = uuid.NewV4().String()
			tablenamesql.Tablename = tablename
			tablenamesql.PutInTime = time.Now().Unix()
			err = db.Table("xsstablelist").Create(&tablenamesql).Error
			if err != nil {
				fmt.Println("操作xsstablelist时出现错误:", err)
				status = 0
				return
			}
			status = 1
			return
		}
		fmt.Println("操作xsstablelist时出现错误:", err)
		status = 0
		return
	}
	// 说明查询到了同名桌子
	status = 2
	return
}
