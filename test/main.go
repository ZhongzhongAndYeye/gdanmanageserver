package main

import "server/utils"

type Test2 struct{
	A string
	B string
}

func main() {
	db := utils.DB
	db.AutoMigrate(Test2{})
}