package main

import "fmt"

func main() {
	fmt.Println("Hello, Sunandan Sarkar, welcome to Go programming language.")
	foo()
	fmt.Println("Something More.")
	//Creating a loop
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("I'm in foo")

}

func bar() {
	fmt.Println("And when we exited")

}

//controll flow
//loop; iterative
//conditional
