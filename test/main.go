package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"

)

func main() {
	// var arr1 = []int{1,2,3}
	// var arr2 = []int{2,1,3}
	pj := "[41,3,35,1,6,37,59,45,24,18,49,39,27,26,45,29,57,56,38,2,17,59,8,27,35,51,11]"
	fmt.Printf("%T", pj)
	fmt.Println(pj)
	pjbyte := []byte(pj)
	fmt.Printf("%T", pjbyte)
	fmt.Println(pjbyte[0])
	// sort.Ints(arr2)
	// fmt.Println(reflect.DeepEqual(arr1,arr2))
	// fmt.Println(arr1)
	// fmt.Println(arr2)
	var x []int
	_ = json.Unmarshal(pjbyte, &x)
	fmt.Printf("%T", x)
	sort.Ints(x)
	x2 := []int{1, 2, 3, 6, 8, 11, 17, 18, 24, 26, 27, 27, 29, 35, 35, 37, 38, 39, 41, 45, 45, 49, 51, 56, 57, 59, 59}
	fmt.Println(reflect.DeepEqual(x,x2))

}
