package main

import "fmt"

func main() {
	var a int = 10
	for a < 20 {
		if a == 15 {
			a = a + 1
			//			a++
			continue
		}
		fmt.Printf("a的值为 %d\n", a)
		a++
	}
}
