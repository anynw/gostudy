package main

import "fmt"

func main() {
	var i int = 5
	fmt.Printf("%d 的阶乘 %d\n", i, Factorial(uint64(i)))
	fmt.Printf("%d 的阶乘 %d\n", i, Factorial2(uint64(i)))

	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t", Fib(i))
	}
	fmt.Println()

}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func Factorial2(n uint64) uint64 {
	if n > 0 {
		return n * Factorial2(n-1)
	}
	return 1
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-2) + Fib(n-1)
}
