package model

import (
	"fmt"
	"reflect"
	"server/utils"
	"sort"
	"time"

	"github.com/goccy/go-json"
	"github.com/satori/go.uuid"
)

func Xsspaijuinsert(data [][]int) (msg string) {
	type Xsspaijudata struct {
		Id       int
		Pjid     string
		Order    int
		HandCard string
	}
	type Xsspjid struct {
		Pjid      string
		PutInTime string
		Frequency float64
		Count     int
	}
	db := utils.DB
	// 检查插入的牌局和限时赛仓库中的牌局是否重复
	// 取出库存牌局的牌局id列表
	type Pjidlist struct {
		Pjid string
	}
	pjidlist := []Pjidlist{}
	err := db.Table("xsspjidlist").Find(&pjidlist).Error
	if err != nil {
		msg = "操作xsspjidlist发生错误..."
		fmt.Println(err)
		return
	}
	// 若是没有库存牌局 则直接插入
	if len(pjidlist) == 0 {
		//插入
		pjid := uuid.NewV4().String()
		for i := 0; i < 4; i++ {
			xsspjdata := Xsspaijudata{}
			xsspjdata.Pjid = pjid
			xsspjdata.Order = i + 1
			a, _ := json.Marshal(data[i])
			xsspjdata.HandCard = string(a)
			err = db.Table("xsspaijudatas").Create(&xsspjdata).Error
			if err != nil {
				msg = "操作xsspaijudatas发生错误..."
				fmt.Println(err)
				return
			}
		}
		xsspjid := Xsspjid{}
		xsspjid.Pjid = pjid
		xsspjid.PutInTime = time.Now().Format("2006-01-02 15:04:05")

		err = db.Table("xsspjidlist").Create(&xsspjid).Error
		if err != nil {
			fmt.Println("err")
			msg = "操作xsspjidlist发生错误..."
			return
		}
		msg = "插入限时赛牌局数据成功！"
		return
	} else {
		// 取出所有的牌局数据
		type Pjdata struct {
			Order    int
			HandCard string
		}
		pjdatas := [][]Pjdata{}
		for _, val := range pjidlist {
			entire := []Pjdata{}
			err = db.Table("xsspaijudatas").Where("Pjid = ?", val.Pjid).Order("`order` asc").Find(&entire).Error
			if err != nil {
				msg = "操作xsspaijudatas出现问题..."
				fmt.Println(err)
				return
			}
			pjdatas = append(pjdatas, entire)
		}
		// 进行比较是否重复
		for _, val := range pjdatas {
			count := 0
			for i := 0; i < 4; i++ {
				valbyte := []byte(val[i].HandCard)
				var valintslice []int
				json.Unmarshal(valbyte, &valintslice)
				sort.Ints(data[i])
				sort.Ints(valintslice)
				if reflect.DeepEqual(data[i], valintslice) {
					count = count + 1
				}
			}
			if count == 4 {
				msg = "库存中已存在此牌局，插入失败..."
				return
			}
		}

		// 若是上面的循环没有return 说明无重复牌局
		pjid := uuid.NewV4().String()
		for i := 0; i < 4; i++ {
			xsspjdata := Xsspaijudata{}
			xsspjdata.Pjid = pjid
			xsspjdata.Order = i + 1
			a, _ := json.Marshal(data[i])
			xsspjdata.HandCard = string(a)
			err = db.Table("xsspaijudatas").Create(&xsspjdata).Error
			if err != nil {
				msg = "操作xsspaijudatas发生错误..."
				fmt.Println(err)
				return
			}
		}
		xsspjid := Xsspjid{}
		xsspjid.Pjid = pjid
		xsspjid.PutInTime = time.Now().Format("2006-01-02 15:04:05")

		err = db.Table("xsspjidlist").Create(&xsspjid).Error
		if err != nil {
			fmt.Println("err")
			msg = "操作xsspjidlist发生错误..."
			return
		}
		msg = "插入限时赛牌局数据成功！"
		return
	}
}
