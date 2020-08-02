package main

import (
	"fmt"
	"time"
)

func main() {

	foo("Number")

	go foo("Goroutine")

	go func(ms string) {
		fmt.Println(ms)
	}("Going")

	time.Sleep(time.Second)
	fmt.Println("Program Done")
}
func foo(f string) {
	for i := 0; i < 3; i++ {
		fmt.Println(f, ":", i)
	}
}
