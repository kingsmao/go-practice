package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	e := 100
	return e
}

//多个返回参数，匿名形式
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 111, 222
}

func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	//给有名称的返回值变量赋值
	r1 = 321
	r2 = 543
	return
	//return  这三种形式均可
	//return 11111,22222
	//return r1,r2
}
func main() {
	c := foo1("abc", 233)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("dddd", 123)
	fmt.Println("ret1 = ", ret1, ",ret2 = ", ret2)

	r1, r2 := foo3("foo3", 23456)
	fmt.Println("r1 = ", r1, " r2 = ", r2)
}
