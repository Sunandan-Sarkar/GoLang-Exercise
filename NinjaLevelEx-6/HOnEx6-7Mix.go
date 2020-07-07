package main

import (
	"fmt"
)

var x int
var y func()

func main() {
	func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}()
	fmt.Println("done")
	f1 := func() {
		for i := 0; i < 9; i++ {
			fmt.Println(i)
		}
	}
	f1()
	fmt.Println("Done!")
	fmt.Printf("%T\n%T\n", x, f1)
	y = f1
	y()
	fmt.Printf("%T\n", y)
}
