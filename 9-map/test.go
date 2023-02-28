package main

import "fmt"

func main() {
	var myMap1 map[string]string
	myMap1 = make(map[string]string, 10) //第一种声明 + 分配空间
	myMap1["one"] = "java"
	myMap1["two"] = "php"
	myMap1["three"] = "go"
	fmt.Println(myMap1)

	myMap2 := make(map[string]string) //第二种，容量可以省去
	myMap2["one"] = "java"
	myMap2["two"] = "php"
	myMap2["three"] = "go"
	fmt.Println(myMap2)

	myMap3 := map[string]string{ //第三种形式
		"one":   "go1",
		"two":   "go2",
		"three": "go3",
	}
	fmt.Println(myMap3)
}
