package model

import (
	"encoding/json"
	"fmt"
	"reflect"
	"server/utils"
	"sort"
)

type Pjdataslice struct {
	Order    int
	HandCard []int
}

func Diyxssyx(tableid string, tablepjid string, data []Pjdataslice) (status int) {
	// fmt.Println(data)
	db := utils.DB
	// 查询桌子下所有的牌局id
	type Tablepjid struct {
		Tablepjid string
	}
	tablepjlist := []Tablepjid{}
	err := db.Table("xsstablepj").Where("tableid = ?", tableid).Find(&tablepjlist).Error
	if err != nil {
		fmt.Println("操作xsstablepj时出现错误...", err)
		status = 0
		return
	}
	// fmt.Println(tablepjlist)

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
	// fmt.Println(tablepjdatas)

	// 遍历桌子下的每一个牌局，检查是否和传入的牌局重复
	for _, val := range tablepjdatas {
		//遍历一局的四个order
		count := 0
		for i := 0; i < 4; i++ {
			handcardbyte := []byte(val[i].HandCard)
			var handcardint []int
			json.Unmarshal(handcardbyte, &handcardint)
			sort.Ints(data[i].HandCard)
			sort.Ints(handcardint)
			if reflect.DeepEqual(data[i].HandCard, handcardint) {
				count = count + 1
			}
		}
		if count == 4 {
			fmt.Println("此桌子上已经存在此牌局")
			status = 2
			return
		}
	}

	// 若是桌子中没有重复的，那么遍历库存牌局查看是否有重复
	// 查询库存牌局中的牌局列表
	type Kcpjid struct {
		Pjid string
	}
	kcpjidlist := []Kcpjid{}
	err = db.Table("xsspjidlist").Find(&kcpjidlist).Error
	if err != nil {
		fmt.Println("操作xsspjidlist时出现错误...")
		status = 0
		return
	}
	// fmt.Println(kcpjidlist)

	// 获取每局库存牌局的详细数据
	kcpjdatas := [][]Tablepjdata{}
	for _, val := range kcpjidlist {
		entire := []Tablepjdata{}
		err = db.Table("xsspaijudatas").Where("pjid = ?", val.Pjid).Order("`order` asc").Find(&entire).Error
		if err != nil {
			fmt.Println("操作xsstablepjdatas时发生错误...", err)
			status = 0
			return
		}
		kcpjdatas = append(kcpjdatas, entire)
	}
	// fmt.Println(kcpjdatas)

	// 遍历库存牌局下的每一个牌局，检查是否和传入的牌局重复
	for _, val := range kcpjdatas {
		count := 0
		for i := 0; i < 4; i++ {
			handcardbyte := []byte(val[i].HandCard)
			var handcardint []int
			json.Unmarshal(handcardbyte, &handcardint)
			sort.Ints(handcardint)
			sort.Ints(data[i].HandCard)
			if reflect.DeepEqual(handcardint, data[i].HandCard) {
				count = count + 1
			}
		}
		if count == 4 {
			fmt.Println("库存牌局中存在重复牌局")
			status = 3
			return
		}
	}

	// 桌子上没有重复的，库存也没有重复的，那么对桌子上改动的那局牌局进行数据更新
	for i := 1; i <= 4; i++ {
		handcardbytecon,_ := json.Marshal(data[i-1].HandCard)
		handcardcon := string(handcardbytecon)
		err = db.Table("xsstablepjdatas").Where("tablepjid = ?",tablepjid).Where("`order` = ?",i).Update("hand_card",handcardcon).Error	
		if err != nil {
			fmt.Println("操作xsstablepjdatas时出现错误...")
			status = 0
			return
		}
	}
	
	status = 1
	return
}
