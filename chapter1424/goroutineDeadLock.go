package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("goroutine deadlock")
	out := make(chan int)
	fmt.Println("sending .....")
	out <- 2
	fmt.Println("sent")
	go f1(out)
}

func f1(in chan int) {
	time.Sleep(10 * 1e9)
	fmt.Println("receiving .....")
	fmt.Println(<-in)
	fmt.Println("received")
}
