package main

import "fmt"

func main() {
	// 全局
	var arr0 = [5]int{1, 2, 3}
	var arr1 = [5]int{1, 2, 3, 4, 5}
	var arr2 = [...]int{1, 2, 3, 4, 5}

	// 局部

	fmt.Println(arr0)
	fmt.Println(arr1)
	fmt.Println(arr2)
}
