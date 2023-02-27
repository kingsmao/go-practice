package main

import "fmt"

func main() {

	//int 默认值0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n",a)

	//给定类型赋值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n",b)

	//不给类型，程序自动推断类型
	var c  = "200"
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n",c)

	//省去var ，只能用在函数体中，不能用在全局变量中
	e:= "100"
	fmt.Printf("type of e = %T\n",e)


}
