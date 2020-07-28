package main

import "fmt"

func main() {
	cp := make(chan int)
	//Send
	go foo(cp)

	//Receive
	 bar(cp)
	fmt.Println("About to exit")
}

//Send
func foo(cp chan<- int) {
	cp <- 84
}

//Receive
func bar(cp <-chan int) {
	fmt.Println(<-cp)
}
