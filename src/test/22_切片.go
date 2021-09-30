package main

import "fmt"

/*
1.Go 语言切片是对数组的抽象。
2.Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型
切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
*/
func main() {
	// 切片初始化
	s := []int{11, 22, 33}
	arr := s[:]
	arr1 := s[0:1]
	arr2 := s[1:2]
	arr3 := s[0:2]
	arr4 := s[1:3]
	arr5 := s[0:3]

	arr6 := s[0:]
	arr7 := s[:3]

	s1 := make([]int, 3, 3)
	s2 := make([]int, 2, 3)

	fmt.Printf("arr = %v\n", arr)
	fmt.Printf("arr1 = %v\n", arr1)
	fmt.Printf("arr2 = %v\n", arr2)
	fmt.Printf("arr3 = %v\n", arr3)
	fmt.Printf("arr4= %v\n", arr4)
	fmt.Printf("arr5= %v\n", arr5)
	fmt.Printf("arr6= %v\n", arr6)
	fmt.Printf("arr7= %v\n", arr7)
	fmt.Printf("s1= %v\n", s1)
	fmt.Printf("s2= %v\n", s2)
}
