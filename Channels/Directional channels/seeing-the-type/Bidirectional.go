package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	cr := make(<-chan int) // Receive
	cs := make(chan<- int) // Send

	fmt.Println("------------------")
	fmt.Printf("C:\t%T\n", c)
	fmt.Printf("CR:\t%T\n", cr)
	fmt.Printf("CS:\t%T\n", cs)

}
