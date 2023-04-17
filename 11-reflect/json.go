package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"acts"`
}

func main() {
	m := Movie{"带带弟弟", 2023, 10, []string{"d1", "d2", "d3"}}
	marshal, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr = %s\n", marshal)
	myMovie := Movie{}
	err = json.Unmarshal(marshal, &myMovie)
	fmt.Println(myMovie)
}
