package model

import (
	"encoding/json"
	"fmt"
	"reflect"
	"server/utils"
	"sort"
)

func Changexssyxorder(tablepjid string, beforeo int, aftero int, tableid string) (status int) {
	db := utils.DB
	// 将序号调换
	err := db.Table("xsstablepjdatas").Where("tablepjid = ?", tablepjid).Where("`order` = ?", aftero).Update("order", 666).Error
	if err != nil {
		fmt.Println("操作xsstablepjdatas时出现错误1...", err)
		status = 0
		return
	}
	err = db.Table("xsstablepjdatas").Where("tablepjid = ?", tablepjid).Where("`order` = ?", beforeo).Update("order", aftero).Error
	if err != nil {
		fmt.Println("操作xsstablepjdatas时出现错误2...", err)
		status = 0
		return
	}
	err = db.Table("xsstablepjdatas").Where("tablepjid = ?", tablepjid).Where("`order` = ?", 666).Update("order", beforeo).Error
	if err != nil {
		fmt.Println("操作xsstablepjdatas时出现错误3...", err)
		status = 0
		return
	}

	// 调换后查询此桌子下是否有相同牌局，若是有则将顺序调换回去并且返回status=2表示牌局重复
	// 查询此桌子下所有的桌子牌局id
	type Tablepj struct {
		Tablepjid string
	}
	tablepjlist := []Tablepj{}
	err = db.Table("xsstablepj").Where("tableid = ?", tableid).Find(&tablepjlist).Error
	if err != nil {
		fmt.Println("操作xsstablepj时发生错误...", err)
		status = 0
		return
	}
	fmt.Println(tablepjlist)
	// 查询此桌子上每个桌子牌局id的具体牌局信息(除了进行修改的这个牌局)
	type Tablepjdata struct {
		Order    string
		HandCard string
	}
	tablepjdatas := [][]Tablepjdata{}
	for _, val := range tablepjlist {
		entire := []Tablepjdata{}
		if val.Tablepjid != tablepjid {
			err = db.Table("xsstablepjdatas").Where("tablepjid = ?", val.Tablepjid).Order("`order` asc").Find(&entire).Error
			if err != nil {
				fmt.Println("操作xsstablepjdatas时发生错误...", err)
				status = 0
				return
			}
			tablepjdatas = append(tablepjdatas, entire)
		}
	}
	// 取出刚刚更改了order顺序的tablepj
	pjdatas := []Tablepjdata{}
	err = db.Table("xsstablepjdatas").Where("tablepjid = ?", tablepjid).Order("`order` asc").Find(&pjdatas).Error
	if err != nil {
		fmt.Println("操作xsstablepjdatas时出现错误...", err)
		status = 0
		return
	}

	//比较更改了顺序后此桌子上是否有重复的牌局
	// 若是桌子上还有其他牌局
	if len(tablepjdatas) != 0 {
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
				json.Unmarshal(pjdatasbyte, &pjdatasintslice)
				json.Unmarshal(valbyte, &valintslice)
				sort.Ints(pjdatasintslice)
				sort.Ints(valintslice)
				fmt.Println(pjdatasintslice)
				fmt.Println(valintslice)
				if reflect.DeepEqual(pjdatasintslice, valintslice) {
					count = count + 1
				}
			}
			// 若是发现和某一局牌局重复
			if count == 4 {
				// 恢复原来的order顺序
				err := db.Table("xsstablepjdatas").Where("tablepjid = ?", tablepjid).Where("`order` = ?", beforeo).Update("order", 666).Error
				if err != nil {
					fmt.Println("操作xsstablepjdatas时出现错误1...", err)
					status = 0
					return
				}
				err = db.Table("xsstablepjdatas").Where("tablepjid = ?", tablepjid).Where("`order` = ?", aftero).Update("order", beforeo).Error
				if err != nil {
					fmt.Println("操作xsstablepjdatas时出现错误2...", err)
					status = 0
					return
				}
				err = db.Table("xsstablepjdatas").Where("tablepjid = ?", tablepjid).Where("`order` = ?", 666).Update("order", aftero).Error
				if err != nil {
					fmt.Println("操作xsstablepjdatas时出现错误3...", err)
					status = 0
					return
				}
				status = 2
				return
			}
		}
	}
	status = 1
	return
}
