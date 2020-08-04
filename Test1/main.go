package main

import "fmt"

func main() {
	func() {
		for i := 0; i < 10; i++ {
			if i%3 == 0 {
				fmt.Println("Buzz")
			} else {
				fmt.Println(i)
			}
		}
	}()
	raw()
}
func raw() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Fuzz")
		} else {
			fmt.Println(i)
		}
	}
}
