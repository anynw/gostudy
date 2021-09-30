package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Printf("a type is %T\n", a)
	fmt.Printf("a = %d ; b = %d\n", a, b)
	// 声明了但是不使用变量会报错
	// var c int

}
