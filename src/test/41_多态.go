package main

import "fmt"

// 指针
type Animal interface {
	Sleep()
	GetType() string
	GetColor() string
}

// 具体的类：Cat
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("cat is sleep")
}

func (this *Cat) GetType() string {
	return "Cat"
}

func (this *Cat) GetColor() string {
	return this.color
}

// -------
// 具体的类：Dog
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("dog is sleep")
}

func (this *Dog) GetType() string {
	return "Dog"
}

func (this *Dog) GetColor() string {
	return this.color
}

func showAnimal(animal Animal) {
	animal.Sleep()
}

func main() {
	var animal Animal
	animal = &Cat{"Green"}

	animal.Sleep()

	animal = &Dog{"Yellow"}

	animal.Sleep()

	cat := &Cat{"Green"}
	dog := &Dog{"Yellow"}

	//封装方法
	showAnimal(cat)
	showAnimal(dog)
}
