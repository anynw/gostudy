package main

import "fmt"

func main() {
	var a = 1
	fmt.Println("a = ", a)

	var b int
	b = 2
	fmt.Println("b = ", b)

	c := 3
	fmt.Println("c = ", c)

	var (
		l int
		m int
		n int
	)
	fmt.Printf("l = %d; m = %d; n = %d\n", l, m, n)

	// 多个变量赋值
	aa, bb, cc := 11, 22, 33
	fmt.Printf("aa = %d; bb = %d; cc = %d\n", aa, bb, cc)

}
