package main

import (
	"fmt"
)

func main() {
	tal()
	tartu("Sundan,I'm from Estonia")
}
func tal() {
	fmt.Println("Hello From Tallinn University")
}
func tartu(t string) {
	fmt.Println("Hello", t)
}

//takes an argument
