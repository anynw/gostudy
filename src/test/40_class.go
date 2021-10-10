package main

import "fmt"

type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(newName string) {
	this.Name = newName
}

func main() {
	hero := Hero{Name: "zhang3", Ad: 100, Level: 2}

	hero.Show()

	name := hero.GetName()
	fmt.Println(name)

	hero.SetName("lisi")
	hero.Show()
}
