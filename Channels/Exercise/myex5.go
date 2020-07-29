package main

import "fmt"

func main() {
	f := make(chan int)
	b := make(chan int)

	go func() {
		for i := 0; i < 5000; i++ {
			f <- i
		}
	}()

	go func() {
		for i := 0; i < 5000; i++ {
			b <- i
		}
	}()
	foo(f, b)
	fmt.Println("About to exit")
}
func foo(f, b chan int) {
	for i := 0; i < 10000; i++ {
		select {
		case v := <-f:
			fmt.Println("From Foo:", v)
		case v := <-b:
			fmt.Println("From Bar:", v)
		}
	}
	fmt.Println("About to exit too")
}
