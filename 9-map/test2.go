package main

import "fmt"

func printMap(cityMap map[string]string) {
	//cityMap为指针传递
	for key, value := range cityMap {
		fmt.Println("key = ", key, " value = ", value)
	}
	cityMap["aaa"] = "vvv"
}

func main() {
	cityMap := make(map[string]string)
	//添加
	cityMap["zh"] = "beijing"
	cityMap["usa"] = "dc"
	cityMap["jp"] = "tk"
	//遍历
	printMap(cityMap)
	fmt.Println("----")
	//修改
	cityMap["zh"] = "dddd"
	printMap(cityMap)
	fmt.Println("----")
	//删除
	delete(cityMap, "usa")
	printMap(cityMap)

}
