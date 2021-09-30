package main

import (
	"fmt"
)

// 通过循环嵌套的方式输出2-100之间的素数
func main() {
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
