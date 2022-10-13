package utils

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	DB = SqlConn()
}

// 连接数据库
func SqlConn() (db *gorm.DB) {
	v, err := ReadYaml("mysql") // 读取配置文件获得数据库信息
	if err != nil {
		log.Fatal(err)
		return
	}
	url := v.GetString("db.url")
	username := v.GetString("db.username")
	password := v.GetString("db.password")
	port := v.GetString("db.port")
	sqlname := v.GetString("db.sqlname")

	str := username + ":" + password + "@(" + url + ":" + port + ")/" + sqlname + "?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(str), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库时出错，错误为:", err)
		return
	}
	fmt.Println("数据库已连接成功...")
	return
}
