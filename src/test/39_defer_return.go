package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer  run ......")
	return 0
}

func returnFunc() int {
	fmt.Println("return  run ......")
	return 1
}

func deferAndReturn() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	deferAndReturn()
}
