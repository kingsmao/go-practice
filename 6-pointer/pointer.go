package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println("a = ", a, " b = ", b)
}
func swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}
