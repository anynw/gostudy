package main

import "fmt"

func main() {
	//var var_name *var-type
	//var ip *int        /* 指向整型*/
	//var fp *float32    /* 指向浮点型 */

	var a int = 20
	var ip *int
	ip = &a
	fmt.Printf("a的变量的地址：%x\n", &a)

	fmt.Printf("ip的变量储存指针地址：%x\n", ip)

	fmt.Printf("*i变量的值：%d\n", *ip)

	var ptr *int
	fmt.Printf("ptr 的值：%x\n", ptr)

	// 判断是否空指针
	if ptr != nil {
		fmt.Println("ptr不是空指针")
	} else {
		fmt.Println("ptr是空指针")
	}
}
