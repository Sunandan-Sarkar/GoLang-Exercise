package main

import "fmt"

func main() {
	fmt.Println(bar()())
	fmt.Print(goo()())
	fmt.Println(hoo()())
}
func bar() func() int {
	return func() int {
		return 451
	}
}

func goo() func() string {
	return func() string {
		return "Limbaroyati is a"
	}
}
func hoo() func() []string {
	return func() []string {
		return []string{"Designer", "Happy", "Healthy"}
	}
}
