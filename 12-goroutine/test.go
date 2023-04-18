package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		fmt.Println("goroutine ...")
		c <- 777
	}()

	num := <-c
	fmt.Println("num = ", num)
}
