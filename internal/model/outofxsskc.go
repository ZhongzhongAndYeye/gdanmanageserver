package model

import (
	"errors"
	"fmt"
	"server/utils"
	"strconv"

	"gorm.io/gorm"
)

func Outofxsskc(tableid string, pjid string) (status int) {
	db := utils.DB
	var xsstablepj struct {
		Tableid string
		Pjid    string
	}
	// 查询指定的桌上中是否已经存在选定牌局
	err := db.Table("xsstablepj").Where("tableid = ?", tableid).Where("pjid = ?", pjid).First(&xsstablepj).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 若是没有查询到结果
			xsstablepj.Tableid = tableid
			xsstablepj.Pjid = pjid
			err = db.Table("xsstablepj").Create(&xsstablepj).Error
			if err != nil {
				fmt.Println("操作xsstablepj时出现错误...", err)
				status = 0
				return
			}

			// 将库存牌局表中的数据复制一份到已选牌局数据表中
			type Orderandhc struct {
				Pjid     string
				Order    int
				HandCard string
			}
			// 查询已选牌局数据表中是否已经有相应的牌局了
			orderandhc := []Orderandhc{}
			err = db.Table("xsstablepjdatas").Where("pjid = ?", pjid).Find(&orderandhc).Error
			if err != nil {
				fmt.Println("操作xsstablepjdatas时出现错误...", err)
				status = 0
				return
			}
			// 若是没有 复制牌局，若是已经存在了则不进行复制操作
			if len(orderandhc) == 0 {
				err = db.Table("xsspaijudatas").Where("pjid = ?", pjid).Find(&orderandhc).Error
				if err != nil {
					fmt.Println("操作xsspaijudatas时出现错误...", err)
					status = 0
					return
				}
				err = db.Table("xsstablepjdatas").Create(&orderandhc).Error
				if err != nil {
					fmt.Println("操作xsstablepjdatas时出现错误...", err)
					status = 0
					return
				}
			}

			// 给指定牌局的出库次数和总操作次数分别加1
			err = db.Table("xsspjidlist").Where("pjid = ?", pjid).Update("count", gorm.Expr("count + ?", 1)).Error
			if err != nil {
				fmt.Println("操作xsspjidlist时出现错误...", err)
				status = 0
				return
			}
			err = db.Table("xssckcount").Where("id = ?", 1).Update("count", gorm.Expr("count + ?", 1)).Error
			if err != nil {
				fmt.Println("操作xssckcount时出现错误...", err)
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
			status = 1
			return
		}
		fmt.Println("操作xsstablepj时发生错误...", err)
		status = 0
		return
	}
	// 查询到了结果
	fmt.Println("同一牌局不可重复出库到同一张桌子！")
	status = 2
	return
}
