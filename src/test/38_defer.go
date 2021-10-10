package main

import "fmt"

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

func main() {
	//	test()
	defer func1()
	defer func2()
	defer func3()
}

func test() {
	// defer 先入后执行
	defer fmt.Println("main:end")
	defer fmt.Println("main:end2")

	fmt.Println("main:go 1")
	fmt.Println("main:go 2")
}
