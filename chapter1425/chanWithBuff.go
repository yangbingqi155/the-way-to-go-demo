package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("channel with buffer")
	//chan without buffer
	//c := make(chan int)
	//chan with buffer
	c := make(chan int, 50)
	go func() {
		time.Sleep(15 * 1e9)
		x := <-c
		fmt.Println("Received ", x)
	}()
	for i := 0; i < 60; i++ {
		fmt.Println("sending ", i)
		c <- i
		fmt.Println("sent ", i)
	}

}
