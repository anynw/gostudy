package main

import "fmt"

/*
如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。
*/
func main() {
	var numbers []int
	printSlice(numbers)

	// 允许追加空切片
	numbers = append(numbers, 0)
	printSlice(numbers)

	// 向切片添加一个元素
	numbers = append(numbers, 1)
	printSlice(numbers)

	// 同时添加多个元素
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	copy(numbers1, numbers)
	printSlice(numbers1)
}

func printSlice(x []int) {
	fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x), x)
}
