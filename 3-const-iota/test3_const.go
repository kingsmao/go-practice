package main

import "fmt"

/*
const (
	AAA = 1
	BBB = 2
	CCC = 3
)
*/

const (
	AAA = 10 * iota //关键字iota，默认值=0，向下自增
	BBB             // BBB = 10 * 递增后的iota = 10 * 1
	CCC             // CCC = 10 * 递增后的iota = 10 * 2
)

//iost只能出现在const中
const (
	a, b = iota + 1, iota + 2 //iota = 0,a = iota + 1,b = iota + 2,a = 1,b = 2
	c, d					  //iota = 1,c = iota + 1,d = iota + 2,c = 2,d = 3
	e, f					  //iota = 2,e = iota + 1,f = iota + 2,e = 3,f = 4
	g, h = iota * 2, iota * 3 //iota = 3,g = iota * 2,h = iota * 3,g = 6,h = 9
	i, k					  //iota = 4,i = iota * 2,k = iota * 3,i = 8,k = 12
)

func main() {
	//const 常量 不可修改
	const length int = 10
	fmt.Println("length = ", length)
	fmt.Println("AAA = ", AAA, ", BBB = ", BBB, "CCC = ", CCC)

}
