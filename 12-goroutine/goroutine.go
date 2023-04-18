package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("newTask i= %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go newTask()
	i := 0
	for {
		i++
		fmt.Printf("main i= %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
