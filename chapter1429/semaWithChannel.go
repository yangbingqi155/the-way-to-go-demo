package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("sema with channel")
	N := 10
	sema := writeNumber(N)
	readNumber(N, sema)
}

func writeNumber(N int) chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < N; i++ {
			time.Sleep(1e9)
			ch <- i * 10
		}
	}()
	return ch
}
func readNumber(N int, ch chan int) {
	for i := 0; i < N; i++ {
		fmt.Println(<-ch)
	}
}
