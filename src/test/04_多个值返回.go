package main

import "fmt"

func main() {
	_, num, strs := test()
	fmt.Println(num, strs)
}

// 多个返回值，应用空白标识符
func test() (int, int, string) {
	a, b, c := 1, 2, "str1"
	return a, b, c
}
