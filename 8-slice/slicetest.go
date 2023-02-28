package main

import (
	"fmt"
)

func main() {
	myArray := []int{1, 2} //[] 动态数组，传递给方法时是指针传递
	for index, value := range myArray {
		fmt.Println("index = ", index)
		fmt.Println("value = ", value)
	}

	mySlice()
}

func mySlice() {
	//声明slice1是一个切片，并且初始化，默认值是1，2，3，长度3
	slice1 := []int{1, 2, 3}
	fmt.Printf("slice1 len = %d,slice1 = %v\n", len(slice1), slice1)

	//声明了一个切片但是没有分配空间，需要配合make使用
	var slice2 []int
	slice2 = make([]int, 3) //开辟3个空间
	fmt.Printf("slice2 len = %d,slice2 = %v\n", len(slice2), slice2)

	//var slice3 []int = make([]int,4)//声明切片同时开辟空间
	slice3 := make([]int, 5) //简写形式
	fmt.Printf("slice3 len = %d,slice3 = %v\n", len(slice3), slice3)
}
