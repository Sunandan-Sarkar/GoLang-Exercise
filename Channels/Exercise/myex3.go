package main

import (
	"fmt"
)

func main() {
	cb := make(chan int)

	go func(cb chan<- int) {
		cb <- 5
	}(cb)

	func(cb <-chan int) {
		fmt.Println(<-cb)
	}(cb)

	fmt.Println("About to EXIT")
}
