package main

import "fmt"

func main() {
	f := too()
	fmt.Println(f())
}
func too() func() int {
	return func() int {
		return 273
	}
}
