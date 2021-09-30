package main

import "fmt"

/*
	传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，
	将不会影响到实际参数。默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数
*/
func main() {
	a := 100
	b := 200
	fmt.Printf("交换前a的值：%d\n", a)
	fmt.Printf("交换前b的值：%d\n", b)

	swap(a, b)

	fmt.Printf("交换后a的值：%d\n", a)
	fmt.Printf("交换后b的值：%d\n", b)
}

func swap(x, y int) int {
	var temp int
	temp = x
	x = y
	y = temp
	return temp
}
