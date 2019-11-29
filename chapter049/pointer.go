package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointer demo")
	i := 5
	fmt.Printf("a integer %d memory address:%p\n", i, &i)
}
