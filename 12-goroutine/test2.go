package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)
	go func() {
		defer fmt.Println("子go程运行结束")
		for i := 0; i < 3; i++ {
			c <- i
			fmt.Println("发送元素i=", i, "len(c)=", len(c), "cap(c)=", +cap(c))
		}
	}()

	time.Sleep(1 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println("num=", num)
	}
}
