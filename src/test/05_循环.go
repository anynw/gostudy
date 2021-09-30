package main

import (
	"fmt"
)

func main() {
	// 死循环
	//	for true {
	//		fmt.Println("true----\n")
	//	}
	//	sum1()
	//	sum2()
	//	sum3()
	//	rangeForEach()
	for1()
}

// 写法1
func sum1() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum)
}

// 写法2
func sum2() {
	sum := 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println("sum = ", sum)
}

// 写法3
func sum3() {
	sum := 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println("sum = ", sum)
}

func rangeForEach() {
	strings := []string{"apple", "pick"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
	/**
	第 0 位 x 的值 = 1
	第 1 位 x 的值 = 2
	第 2 位 x 的值 = 3
	第 3 位 x 的值 = 5
	第 4 位 x 的值 = 0
	第 5 位 x 的值 = 0
	*/
	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}
}

// 通过循环嵌套的方式输出2-100之间的素数
func for1() {
	// 定义局部变量
	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				// 如果发现因子，则不是素数
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d 是素数\n", i)
		}
	}
}
