package main

import "fmt"

// 定义结构体
type Circle struct {
	radius float64
}

func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
