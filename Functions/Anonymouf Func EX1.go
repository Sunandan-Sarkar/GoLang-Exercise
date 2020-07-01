package main

import "fmt"

func main() {
	func(x int) {
		fmt.Println("My Dog price is in euro:", x)
	}(5566)
	func(y string) {
		fmt.Println("My dog name is:", y)
	}("Robin")
	func(p []string) {
		fmt.Println("My all animal", p)
	}([]string{"cat", "dog", "deer"})
}
