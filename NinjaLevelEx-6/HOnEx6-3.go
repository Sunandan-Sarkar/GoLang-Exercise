package main

import (
	"fmt"
)

func main() {
	defer foo()
	bar()

}
func foo() {
	fmt.Println("I run 1st")
}
func bar() {
	fmt.Println("I run second")

}
