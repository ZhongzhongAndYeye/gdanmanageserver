package main

import "fmt"

func main() {
	// 准备两幅牌 数值模16即为牌值 1-13对应a-k 1-13方块 17-29梅花 33-45红桃 49-61黑桃 65小王 66大王
	twoFullCards := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
		49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
		65, 66,
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
		49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
		65, 66,
	}
	fmt.Println(twoFullCards[0])

	// 序号对应的牌型 共26种
	order_number := make([][]int, 26)
	//顺序 四王 八炸 七炸 六炸 同花顺 五炸 四炸 三张 对子 单牌
	order_number[0] = []int{0, 0, 0, 0, 2, 2, 0, 2, 3}
	order_number[1] = []int{0, 0, 0, 0, 2, 1, 1, 3, 2}
	order_number[2] = []int{0, 0, 0, 0, 3, 0, 1, 3, 2}
	order_number[3] = []int{0, 0, 0, 0, 3, 1, 0, 2, 3}
	order_number[4] = []int{0, 0, 0, 0, 4, 0, 0, 2, 3}
	order_number[5] = []int{0, 0, 0, 1, 1, 1, 1, 2, 3}
	order_number[6] = []int{0, 0, 0, 1, 2, 0, 1, 2, 3}
	order_number[7] = []int{0, 0, 0, 1, 2, 1, 0, 2, 2}
	order_number[8] = []int{0, 0, 0, 1, 0, 3, 0, 2, 2}
	order_number[9] = []int{0, 0, 1, 0, 1, 2, 0, 2, 1}
	order_number[10] = []int{0, 0, 1, 0, 1, 1, 1, 2, 2}
	order_number[11] = []int{0, 0, 1, 0, 2, 1, 0, 2, 1}
	order_number[12] = []int{0, 0, 1, 0, 2, 0, 1, 2, 2}
	order_number[13] = []int{0, 0, 0, 2, 2, 0, 0, 2, 1}
	order_number[14] = []int{0, 0, 0, 3, 1, 0, 0, 1, 2}
	order_number[15] = []int{0, 0, 0, 1, 3, 0, 0, 2, 2}
	order_number[16] = []int{0, 3, 0, 0, 0, 0, 0, 0, 3}
	order_number[17] = []int{0, 2, 0, 0, 0, 1, 1, 0, 2}
	order_number[18] = []int{0, 2, 1, 0, 0, 0, 0, 1, 2}
	order_number[19] = []int{0, 1, 0, 2, 0, 1, 0, 0, 2}
	order_number[20] = []int{0, 0, 0, 0, 3, 0, 2, 1, 2}
	order_number[21] = []int{1, 0, 0, 0, 2, 2, 0, 1, 1}
	order_number[22] = []int{1, 0, 0, 0, 3, 1, 0, 1, 1}
	order_number[23] = []int{1, 2, 0, 0, 0, 0, 1, 1, 1}
	order_number[24] = []int{1, 0, 0, 1, 2, 0, 1, 1, 1}
	order_number[25] = []int{1, 0, 0, 2, 1, 0, 1, 0, 2}

	// 组合对应的四个牌型 共37种 对应的组合是order_number的下标
	combos := make([][]int, 37)
	combos[0] = []int{0, 0, 1, 2}
	combos[1] = []int{1, 1, 1, 1}
	combos[2] = []int{2, 2, 2, 2}
	combos[3] = []int{3, 1, 2, 1}
	combos[4] = []int{0, 0, 1, 1}
	combos[5] = []int{1, 1, 0, 0}
	combos[6] = []int{1, 1, 2, 2}
	combos[7] = []int{2, 2, 1, 2}
	combos[8] = []int{2, 0, 2, 3}
	combos[9] = []int{3, 2, 0, 2}
	combos[10] = []int{0, 1, 2, 3}
	combos[11] = []int{3, 2, 1, 0}
	combos[12] = []int{4, 2, 3, 1}
	combos[13] = []int{20, 20, 20, 20}
	combos[14] = []int{6, 5, 5, 6}
	combos[15] = []int{6, 6, 6, 6}
	combos[16] = []int{7, 6, 7, 6}
	combos[17] = []int{5, 6, 5, 7}
	combos[18] = []int{7, 2, 2, 1}
	combos[19] = []int{7, 6, 3, 2}
	combos[20] = []int{3, 2, 7, 6}
	combos[21] = []int{11, 13, 15, 12}
	combos[22] = []int{13, 15, 13, 11}
	combos[23] = []int{20, 2, 4, 2}
	combos[24] = []int{9, 10, 9, 11}
	combos[25] = []int{10, 9, 10, 9}
	combos[26] = []int{11, 12, 11, 12}
	combos[27] = []int{13, 14, 15, 15}
	combos[26] = []int{14, 15, 13, 15}
	combos[29] = []int{16, 17, 18, 19}
	combos[30] = []int{17, 18, 19, 16}
	combos[31] = []int{18, 19, 16, 17}
	combos[32] = []int{11, 1, 2, 3}
	combos[33] = []int{22, 3, 6, 6}
	combos[34] = []int{23, 17, 18, 19}
	combos[35] = []int{24, 6, 7, 5}
	combos[36] = []int{25, 13, 13, 15}

	

}
