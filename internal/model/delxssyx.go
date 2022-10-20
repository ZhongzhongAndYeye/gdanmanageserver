package model

import (
	"errors"
	"fmt"
	"server/utils"
	"strconv"

	"gorm.io/gorm"
)

func Delxssyx(tableid string, pjid string) (status int) {
	db := utils.DB
	type Tablepj struct {
		Id      int
		Tableid string
		Pjid    string
	}
	// 删除指定桌子上的指定牌局记录
	err := db.Table("xsstablepj").Where("tableid = ?", tableid).Where("pjid = ?", pjid).Delete(&Tablepj{}).Error
	if err != nil {
		fmt.Println("操作xsstablepj表时出现问题", err)
		status = 0
		return
	}
	// 查找其他桌子上是否还有此牌局
	tablepjs := []Tablepj{}
	err = db.Table("xsstablepj").Where("pjid = ?", pjid).Find(&tablepjs).Error
	if err != nil {
		fmt.Println("操作xsstablepj表时出现问题", err)
		status = 0
		return
	}
	// 若是其他桌子上没有此牌局 在已选牌局数据表上删除此牌局
	if len(tablepjs) == 0 {
		type Pjdata struct {
			Id       int
			Pjid     string
			Order    int
			HandCard string
		}
		err = db.Table("xsstablepjdatas").Where("pjid = ?", pjid).Delete(Pjdata{}).Error
		if err != nil {
			fmt.Println("操作xsstablepjdatas表时出现问题", err)
			status = 0
			return
		}
	}
	// 将操作总数减1
	err = db.Table("xssckcount").Where("id = ?", 1).Update("count", gorm.Expr("count - ?", 1)).Error
	if err != nil {
		fmt.Println("操作xssckcount表时出现问题", err)
		status = 0
		return
	}
	// 查询限时赛库存牌局表中的指定牌局是否被删除
	var kcpj struct {
		Pjid string
	}
	err = db.Table("xsspjidlist").Where("pjid = ?", pjid).First(&kcpj).Error
	if err != nil {
		// 若是没有查询到 表示已经被删除了，直接重新计算所有牌局的采用率
		if errors.Is(err, gorm.ErrRecordNotFound) {
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
		fmt.Println("操作xsspjidlist表时出现问题", err)
		status = 0
		return
	}
	// 若是此牌局没有从库存中删除
	// 给pjid对应的操作数减1
	err = db.Table("xsspjidlist").Where("pjid = ?", pjid).Update("count", gorm.Expr("count - ?", 1)).Error
	if err != nil {
		fmt.Println("操作xsspjidlist时出现错误...", err)
		status = 0
		return
	}
	// 查询操作后的操作数是否为0 若是为0直接置空采用率（因为之后的遍历针对的时操作数大于0的牌局）
	var count struct {
		Count int
	}
	err = db.Table("xsspjidlist").Where("pjid = ?", pjid).First(&count).Error
	if err != nil {
		fmt.Println("操作xsspjidlist时出现错误...", err)
		status = 0
		return
	}
	if count.Count == 0 {
		err = db.Table("xsspjidlist").Where("pjid = ?", pjid).Update("frequency", 0).Error
		if err != nil {
			fmt.Println("操作xsspjidlist时出现错误...", err)
			status = 0
			return
		}
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
