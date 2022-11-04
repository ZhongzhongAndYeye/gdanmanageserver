package model

import (
	"encoding/json"
	"fmt"
	"reflect"
	"server/utils"
	"sort"
	"strconv"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func Outofxsskc(tableid string, pjid string) (status int) {
	db := utils.DB

	// 查询此桌子下所有的桌子牌局id
	type Tablepj struct {
		Tablepjid string
	}
	tablepjlist := []Tablepj{}
	err := db.Table("xsstablepj").Where("tableid = ?", tableid).Find(&tablepjlist).Error
	if err != nil {
		fmt.Println("操作xsstablepj时发生错误...", err)
		status = 0
		return
	}

	// 查询此桌子上每个桌子牌局id的具体牌局信息
	type Tablepjdata struct {
		Order    int
		HandCard string
	}

	tablepjdatas := [][]Tablepjdata{}
	for _, val := range tablepjlist {
		entire := []Tablepjdata{}
		err = db.Table("xsstablepjdatas").Where("tablepjid = ?", val.Tablepjid).Order("`order` asc").Find(&entire).Error
		if err != nil {
			fmt.Println("操作xsstablepjdatas时发生错误...", err)
			status = 0
			return
		}
		tablepjdatas = append(tablepjdatas, entire)
	}

	// 遍历桌子上的每一局牌局，若是有完全相同的则出库失败，否则才可以出库
	// 取出对应的牌局数据
	type Pjdata struct {
		Order    int
		HandCard string
	}
	pjdatas := []Pjdata{}
	err = db.Table("xsspaijudatas").Where("pjid = ?", pjid).Order("`order` asc").Find(&pjdatas).Error
	if err != nil {
		fmt.Println("操作xsspaijudatas时发生错误...", err)
		status = 0
		return
	}
	// 若是没有查询到数据 直接插入
	if len(tablepjdatas) == 0 {
		status = insert(tableid, pjid)
		return
	}
	// val是一个数组表示一整局手牌，里面有四个对象
	for _, val := range tablepjdatas {
		// 遍历四局手牌
		count := 0
		for i := 0; i < 4; i++ {
			// 解析数据库中的手牌 string-->[]int并且排序 才能比较其中的元素是否相同
			pjdatasbyte := []byte(pjdatas[i].HandCard)
			valbyte := []byte(val[i].HandCard)
			var pjdatasintslice []int
			var valintslice []int
			json.Unmarshal(pjdatasbyte,&pjdatasintslice)
			json.Unmarshal(valbyte,&valintslice)
			sort.Ints(pjdatasintslice)
			sort.Ints(valintslice)
			fmt.Println(pjdatasintslice)
			fmt.Println(valintslice)
			if reflect.DeepEqual(pjdatasintslice,valintslice) {
				count = count + 1
			}
		}
		if count == 4 {
			status = 2
			return
		}
	}
	// 将操作总数加1
	err = db.Table("xssckcount").Where("id = ?", 1).Update("count", gorm.Expr("count + ?", 1)).Error
	if err != nil {
		fmt.Println("操作xssckcount表时出现问题", err)
		status = 0
		return
	}
	// 将牌局操作数加1
	err = db.Table("xsspjidlist").Where("pjid = ?", pjid).Update("count", gorm.Expr("count + ?", 1)).Error
	if err != nil {
		fmt.Println("操作xsspjidlist表时出现问题", err)
		status = 0
		return
	}

	// 取出总操作次数
	var xssckcount struct {
		Count int
	}
	err = db.Table("xssckcount").Select("Count").Where("id = ?", 1).Find(&xssckcount).Error
	if err != nil {
		fmt.Println("操作xsspjidlist表时发生错误...", err)
		status = 0
		return
	}
	// 获取所有操作过的牌局的pjid
	type Con struct {
		Pjid string
	}
	con := []Con{}
	err = db.Table("xsspjidlist").Where("count > ?", 0).Find(&con).Error
	if err != nil {
		fmt.Println("操作xsspjidlist时出现错误...", err)
		status = 0
		return
	}
	// 更新所有出库过的牌局的采用率
	for _, v := range con {
		var xsspjinfo struct {
			Count int
		}
		err = db.Table("xsspjidlist").Select("Count").Where("pjid = ?", v.Pjid).Find(&xsspjinfo).Error
		if err != nil {
			fmt.Println("操作xsspjidlist表时发生错误...", err)
			status = 0
			return
		}
		var frequency float64
		frequency = float64(xsspjinfo.Count) / float64(xssckcount.Count)
		frequency, err = strconv.ParseFloat(fmt.Sprintf("%.4f", frequency), 64)
		if err != nil {
			fmt.Println("保留小数时出现错误...", err)
			status = 0
			return
		}
		err = db.Table("xsspjidlist").Where("pjid = ?", v.Pjid).Update("frequency", frequency).Error
		if err != nil {
			fmt.Println("操作xsspjidlist表时发生错误...", err)
			status = 0
			return
		}
	}
	// 在桌子上生成唯一的桌子牌局id
	var tablepj struct {
		Tableid   string
		Tablepjid string
	}
	tablepj.Tableid = tableid
	tablepj.Tablepjid = uuid.NewV4().String()
	err = db.Table("xsstablepj").Create(&tablepj).Error
	if err != nil {
		fmt.Println("操作xsstablepj时发生错误...", err)
		status = 0
		return
	}

	// 桌子牌局信息插入牌局数据
	type Tablepjdata2 struct {
		Tablepjid string
		Order     int
		HandCard  string
	}
	tablepjdatas2 := make([]Tablepjdata2, 4)
	for i := 0; i < 4; i++ {
		tablepjdatas2[i].Tablepjid = tablepj.Tablepjid
		tablepjdatas2[i].Order = pjdatas[i].Order
		tablepjdatas2[i].HandCard = pjdatas[i].HandCard
	}
	err = db.Table("xsstablepjdatas").Create(&tablepjdatas2).Error
	if err != nil {
		fmt.Println("操作xsstablepjdatas时发生错误...", err)
		status = 0
		return
	}
	status = 1
	return
}





func insert(tableid string, pjid string) (status int) {
	db := utils.DB
	// 取出对应的牌局数据
	type Pjdata struct {
		Order    int
		HandCard string
	}
	pjdatas := []Pjdata{}
	err := db.Table("xsspaijudatas").Where("pjid = ?", pjid).Find(&pjdatas).Error
	if err != nil {
		fmt.Println("操作xsspaijudatas时发生错误...", err)
		status = 0
		return
	}

	// 将操作总数加1
	err = db.Table("xssckcount").Where("id = ?", 1).Update("count", gorm.Expr("count + ?", 1)).Error
	if err != nil {
		fmt.Println("操作xssckcount表时出现问题", err)
		status = 0
		return
	}
	// 将牌局操作数加1
	err = db.Table("xsspjidlist").Where("pjid = ?", pjid).Update("count", gorm.Expr("count + ?", 1)).Error
	if err != nil {
		fmt.Println("操作xsspjidlist表时出现问题", err)
		status = 0
		return
	}

	// 取出总操作次数
	var xssckcount struct {
		Count int
	}
	err = db.Table("xssckcount").Select("Count").Where("id = ?", 1).Find(&xssckcount).Error
	if err != nil {
		fmt.Println("操作xsspjidlist表时发生错误...", err)
		status = 0
		return
	}
	// 获取所有操作过的牌局的pjid
	type Con struct {
		Pjid string
	}
	con := []Con{}
	err = db.Table("xsspjidlist").Where("count > ?", 0).Find(&con).Error
	if err != nil {
		fmt.Println("操作xsspjidlist时出现错误...", err)
		status = 0
		return
	}
	// 更新所有出库过的牌局的采用率
	for _, v := range con {
		var xsspjinfo struct {
			Count int
		}
		err = db.Table("xsspjidlist").Select("Count").Where("pjid = ?", v.Pjid).Find(&xsspjinfo).Error
		if err != nil {
			fmt.Println("操作xsspjidlist表时发生错误...", err)
			status = 0
			return
		}
		var frequency float64
		frequency = float64(xsspjinfo.Count) / float64(xssckcount.Count)
		frequency, err = strconv.ParseFloat(fmt.Sprintf("%.4f", frequency), 64)
		if err != nil {
			fmt.Println("保留小数时出现错误...", err)
			status = 0
			return
		}
		err = db.Table("xsspjidlist").Where("pjid = ?", v.Pjid).Update("frequency", frequency).Error
		if err != nil {
			fmt.Println("操作xsspjidlist表时发生错误...", err)
			status = 0
			return
		}
	}
	// 在桌子上生成唯一的桌子牌局id
	var tablepj struct {
		Tableid   string
		Tablepjid string
	}
	tablepj.Tableid = tableid
	tablepj.Tablepjid = uuid.NewV4().String()
	err = db.Table("xsstablepj").Create(&tablepj).Error
	if err != nil {
		fmt.Println("操作xsstablepj时发生错误...", err)
		status = 0
		return
	}

	// 桌子牌局信息插入牌局数据
	type Tablepjdata struct {
		Tablepjid string
		Order     int
		HandCard  string
	}
	tablepjdatas := make([]Tablepjdata, 4)
	for i := 0; i < 4; i++ {
		tablepjdatas[i].Tablepjid = tablepj.Tablepjid
		tablepjdatas[i].Order = pjdatas[i].Order
		tablepjdatas[i].HandCard = pjdatas[i].HandCard
	}
	err = db.Table("xsstablepjdatas").Create(&tablepjdatas).Error
	if err != nil {
		fmt.Println("操作xsstablepjdatas时发生错误...", err)
		status = 0
		return
	}
	status = 1
	return
}
