package main

import (
	_ "go-practice/5-init/lib1" //_ 匿名导包，可以导了不用，但是可以使用里面的init函数
	//myLib2 "go-practice/5-init/lib2" //默认名字是lib2，myLib2起别名
	. "go-practice/5-init/lib2" //把lib2导入本类中，类似java的静态导包，不推荐使用，容易跟别的包冲突
)

func main() {
	//lib1.Lib1Test()//不使用，但是要用lib1的init函数
	//myLib2.Lib2Test()//别名
	Lib2Test() //.导包
}
