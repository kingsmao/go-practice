package main

import "fmt"

type AnimalIf interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	Color string
}

func (this *Cat) Sleep() {
	fmt.Println("cat is sleeping")
}

func (this *Cat) GetColor() string {
	return this.Color
}
func (this *Cat) GetType() string {
	return "cat"
}

type Dog struct {
	Color string
}

func (this *Dog) Sleep() {
	fmt.Println("dog is sleeping")
}

func (this *Dog) GetColor() string {
	return this.Color
}
func (this *Dog) GetType() string {
	return "dog"
}

func Show(animal AnimalIf) {
	animal.Sleep()
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("type = ", animal.GetType())
}

func main() {

	cat := Cat{"green"}
	dog := Dog{"red"}

	Show(&cat)
	Show(&dog)

}
