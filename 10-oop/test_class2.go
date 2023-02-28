package main

import "fmt"

type Human struct {
	name string
	sex  int
}

func (this *Human) Eat() {
	fmt.Println("Human Eat ....")
}

func (this *Human) Walk() {
	fmt.Println("Human Walk ....")
}

type Superman struct {
	Human //继承Human
	level int
}

//重写了父类的eat方法
func (this *Superman) Eat() {
	fmt.Println("Superman Eat ....")
}

func (this *Superman) Fly() {
	fmt.Println("Superman fly ....")
}
func main() {
	h1 := Human{"zhangsan", 1}
	h1.Eat()
	h1.Walk()
	fmt.Println("--------")
	//s1:=Superman{h1,1}
	//s1:=Superman{Human{"lisi",1},2333}
	var s1 Superman
	s1.sex = 1
	s1.name = "wangwu"
	s1.level = 10086
	s1.Fly()
	s1.Walk()
	s1.Eat()
}
