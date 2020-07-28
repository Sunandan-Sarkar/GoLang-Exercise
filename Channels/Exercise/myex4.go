package main

import (
	"fmt"
)

func main() {
	fuzz := make(chan int)
	buzz := make(chan int)
	quit := make(chan int)

	go send(fuzz, buzz, quit)
	receive(fuzz, buzz, quit)
}
func send(f, b, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			f <- i
		} else {
			b <- i
		}
	}
	q <- 0
}

func receive(f, b, q <-chan int) {
	for {
		select {
		case v := <-f:
			fmt.Println("From fuzz:", v)
		case v := <-b:
			fmt.Println("From buzz:", v)
		case v := <-q:
			fmt.Println("From quit:", v)
			return
		}

	}
}
