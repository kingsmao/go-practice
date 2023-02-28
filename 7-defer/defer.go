package main

import "fmt"

func main() {
	fmt.Println("main hello 1")
	defer fmt.Println("main end1") //defer关键字，最后执行，多个defer压栈，出栈的顺序，后压入的先出栈，先执行
	defer fmt.Println("main end2")
	fmt.Println("main hello 2")
}
