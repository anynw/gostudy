package main

import "fmt"

type Book struct {
	name string
}

func myFunc(arg interface{}) {
	fmt.Println("myFunc is called ......")
	fmt.Println(arg)

	// 断言
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string type")
	} else {
		fmt.Printf("arg is a string type, value = %v\n", value)
	}
}

func main() {
	book := Book{"GoLang"}

	myFunc(book)
	myFunc(100)
	myFunc("abc")
	myFunc(3.14)
}
