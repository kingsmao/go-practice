package lib1

import "fmt"

func init() { //init函数在外部导包的时候先执行
	fmt.Println("lib1 init...")
}

func Lib1Test() {
	fmt.Println("Lib1 test...")
}
