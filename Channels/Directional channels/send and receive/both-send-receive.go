package main

import "fmt"

func main() {
	ck := make(chan int)
	go func() { ck <- 777 }()
	fmt.Println(<-ck)
}
