package model

import (
	"server/utils"
	"time"
)

type Xssdata struct {
	Pjid      string
	PutInTime time.Time
	Frequency float32
	Count     int
}
type Pjdata struct {
	Order    int
	HandCard string // json字符串
}
type Xssdataall struct {
	Pjid      string
	PutInTime string
	Frequency float32
	Count     int
	Pjdata    []Pjdata
}

func Getxsskc() (msg string, rspxssdataall []Xssdataall) {
	db := utils.DB
	xssdata := []Xssdata{}
	pjdata := []Pjdata{}
	err := db.Table("xsspjidlist").Find(&xssdata).Error
	if err != nil{
		msg = "操作xsspjidlist发生错误"
		return
	}
	num := len(xssdata)
	xssdataall := make([]Xssdataall, num)
	for i, v := range xssdata {
		err := db.Table("xsspaijudatas").Where("pjid = ?", v.Pjid).Find(&pjdata).Error
		if err != nil{
			msg = "操作xsspaijudatas发生错误"
			return 
		}
		xssdataall[i].Pjid = v.Pjid
		xssdataall[i].PutInTime = v.PutInTime.Format("2006-01-02 15:04:05")
		xssdataall[i].Frequency = v.Frequency
		xssdataall[i].Count = v.Count
		xssdataall[i].Pjdata = pjdata
	}
	msg = "查询所有限时赛数据成功！"
	rspxssdataall = xssdataall

	return
}
