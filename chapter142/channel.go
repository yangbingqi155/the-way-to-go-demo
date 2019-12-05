package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("channel")
	//chan without buffer
	c := make(chan int)
	//chan with buffer
	//c := make(chan int, 50)
	go func() {
		time.Sleep(15 * 1e9)
		x := <-c
		fmt.Println("Received ", x)
	}()
	fmt.Println("sending ", 10)
	c <- 10
	fmt.Println("sent ", 10)
}
