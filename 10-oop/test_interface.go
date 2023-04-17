package main

import "fmt"

func myFun(arg interface{}) {
	fmt.Println("myFun is called ...")
	fmt.Println(arg)

	_, ok := arg.(string)
	if ok {
		fmt.Println("arg is right type")
	}
}

type Dook struct {
	name string
}

func main() {
	myFun(Dook{"dd"})
	myFun("123")
	myFun(123)
}
