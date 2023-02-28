package main

import "fmt"

func testDefer() int {
	fmt.Println("called defer...")
	return 0
}

func testReture() int {
	fmt.Println("called return...")
	return 1
}

//return要先于defer执行，defer的返回值不会被采纳
func test() int {
	defer testDefer()
	return testReture()
}

func main() {
	a := test()
	fmt.Println("final a = ", a)

}
