package main

import "fmt"

func main() {
	// 使用range求slice的和
	// 1.数组上使用range将传入index和value两个变量，"_"省略了index
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum = ", sum)

	// 2.使用index
	for index, num := range nums {
		if num == 4 {
			fmt.Println("index = ", index)
		}
	}

	// 3.range在map上的使用
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// 4.range用来枚举Unicode字符串，第一个参数字符的索引，第二个是字符本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
