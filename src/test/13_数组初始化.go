package main

import "fmt"

func main() {
	// 1.默认值
	//var balance = [5]float32{}
	// 2.数组长度不确定，使用... 代替，编译器自行判断
	//	balance := [...]float32{1.2, 3.2}
	// 3.指定下标初始化元素
	balance := [...]int{1: 2, 5: 1}
	fmt.Println("数组：", balance)

	// 初始化二维数组，3行4列
	a := [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /* 第三行索引为 2 */
	}

}
