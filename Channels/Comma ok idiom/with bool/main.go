package main

import (
	"fmt"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(eve, odd, quit)

	receive(eve, odd, quit)
	fmt.Println("About to exit")
}
func send(eve, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}
func receive(eve, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-eve:
			fmt.Println("From even:", v)
		case v := <-odd:
			fmt.Println("From odd:", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("From comma okay:", i, ok)
				return
			} else {
				fmt.Println("From comma okay:", i)
			}
		}
	}
}
