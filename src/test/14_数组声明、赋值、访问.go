package main

import "fmt"

func main() {
	// 初始化一个长度为10的数组
	var n [10]int
	// 定义变量
	var i, j int
	for i = 0; i < 10; i++ {
		// 设置元素的值
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("j[ %d ]的值为 %d\n", j, n[j])
	}
}
