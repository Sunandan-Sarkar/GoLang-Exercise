package main

import (
	"fmt"
)

func main() {
	cv := make(<-chan int)
	go func() {
		cv <- 44
	}()
	fmt.Println(<-cv)

	fmt.Println("---------------")
	fmt.Printf("%T\t", cv)
	//error: “invalid operation: cr <- 42 (send to receive-only type <-chan int)”
}
