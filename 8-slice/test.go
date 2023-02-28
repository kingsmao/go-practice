package main

import (
	"fmt"
)

func main() {
	var myArray [10]int
	myArray2 := [10]int{1, 2, 3, 4, 5}
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}
	for y := 0; y < len(myArray2); y++ {
		fmt.Println(myArray2[y])
	}
	for a, b := range myArray2 {
		fmt.Println("a = ", a)
		fmt.Println("b = ", b)
	}
}
