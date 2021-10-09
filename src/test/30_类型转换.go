package main

import "fmt"

// go不支持隐式转换类型
func main() {
	var sum int = 17
	var count int = 5
	var mean float32
	mean = float32(sum) / float32(count)
	fmt.Printf("mean的值为：%f\n", mean)

	var a int64 = 3
	var b int32
	// 报错：cannot use a (type int64) as type int32 in assignmen
	//	b = a
	b = int32(a)

	fmt.Printf("b的值为：%v\n", b)
}
