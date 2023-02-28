package main

import "fmt"

type Hero struct {
	Name string
	age  int
}

func (this *Hero) GetName() {
	fmt.Println("name = ", this.Name)
}

func (this *Hero) setName(newName string) {
	this.Name = newName
}

func main() {
	hero := Hero{Name: "dddd", age: 1}
	fmt.Println(hero)
	hero.GetName()
	hero.setName("lisi")
	fmt.Println(hero)
}
