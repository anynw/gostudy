package main

import "fmt"

func main() {
	const a = 1
	const b = 2
	const c = 3
	fmt.Println(b)
	const (
		a1 = 1
		b1 = 2
		c1 = 3
	)
	fmt.Printf("a1 = %d; b1 = %d; c1 = %d;", a1, b1, c1)
}
