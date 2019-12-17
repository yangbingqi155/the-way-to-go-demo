package main

import (
	"fmt"
)

//Empty empty interface for sem
type Empty interface{}

var empty Empty

//N data length
const N = 50

func main() {
	fmt.Println("goroutine notice")
	data := make([]float64, N)
	res := make([]float64, N)
	sem := make(chan Empty, N)
	for i := 0; i < N; i++ {
		data[i] = float64(i)
	}
	for i, ix := range data {
		go func(i int, ix float64) {
			res[i] = ix * float64(5)
			sem <- empty
		}(i, ix)
	}
	for i := 0; i < N; i++ {
		<-sem
	}
	for i := 0; i < N; i++ {
		fmt.Printf("Index:%d,Value:%f\n", i, res[i])
	}
}
