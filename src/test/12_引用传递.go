package main

import "fmt"

/*
	引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
	引用传递指针参数传递到函数内，以下是交换函数 swap() 使用了引用传递
*/
func main() {
	a := 100
	b := 200
	fmt.Printf("交换前a的值：%d\n", a)
	fmt.Printf("交换前b的值：%d\n", b)

	swap2(&a, &b)

	fmt.Printf("交换后a的值：%d\n", a)
	fmt.Printf("交换后b的值：%d\n", b)
}

func swap2(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
