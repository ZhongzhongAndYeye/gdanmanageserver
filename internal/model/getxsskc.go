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
	Order    int    `json:"order"`
	HandCard string `json:"handcard"` // json字符串
}
type Xssdataall struct {
	Pjid      string   `json:"pjid"`
	PutInTime string   `json:"putintime"`
	Frequency float32  `json:"frequency"`
	Count     int      `json:"count"`
	Pjdata    []Pjdata `json:"pjdata"`
}

func Getxsskc() (msg string, xssdataall []Xssdataall) {
	db := utils.DB
	xssdata := []Xssdata{}
	pjdata := []Pjdata{}
	err := db.Table("xsspjidlist").Order("put_in_time desc").Find(&xssdata).Error
	if err != nil {
		msg = "操作xsspjidlist发生错误"
		return
	}
	num := len(xssdata)
	xssdataall = make([]Xssdataall, num)
	for i, v := range xssdata {
		err := db.Table("xsspaijudatas").Where("pjid = ?", v.Pjid).Find(&pjdata).Error
		if err != nil {
			msg = "操作xsspaijudatas发生错误"
			return
		}
		xssdataall[i].Pjid = v.Pjid
		xssdataall[i].PutInTime = v.PutInTime.Format("2006-01-02 15:04:05")
		xssdataall[i].Frequency = v.Frequency
		xssdataall[i].Count = v.Count
		for j := 0; j <= 3; j++ { 
			xssdataall[i].Pjdata = append(xssdataall[i].Pjdata, pjdata[j]) // 此处必须append，直接赋值会导致全部指向最后一次循环的pjdata
		}
	}
	msg = "查询所有限时赛数据成功"
	return
}
