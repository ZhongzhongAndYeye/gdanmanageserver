package model

import (
	"fmt"
	"server/utils"
	"time"

	"github.com/goccy/go-json"
	"github.com/satori/go.uuid"
)

func Xsspaijuinsert(data [][]int) (msg string) {
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
	pjid := uuid.NewV4().String()
	db := utils.DB
	for i := 0; i < 4; i++ {
		xsspjdata := Xsspaijudata{}	
		xsspjdata.Pjid = pjid
		xsspjdata.Order = i + 1
		a, _ := json.Marshal(data[i])
		xsspjdata.HandCard = string(a)
		err := db.Table("xsspaijudatas").Create(&xsspjdata).Error
		if err != nil {
			msg = "操作xsspaijudatas发生错误..."
			fmt.Println(err)
			return 
		}
	}
	xsspid := Xsspjid{}
	xsspid.Pjid = pjid
	xsspid.PutInTime = time.Now().Format("2006-01-02 15:04:05")
	err := db.Table("xsspjidlist").Create(&xsspid).Error
	if err != nil{
		fmt.Println("err")
		msg = "操作xsspjidlist发生错误..."
		return
	}
	msg = "插入限时赛牌局数据成功！"
	return
}
