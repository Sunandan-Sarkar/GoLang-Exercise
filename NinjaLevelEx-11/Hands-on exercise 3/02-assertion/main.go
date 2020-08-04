package main

import "fmt"

type customErr struct {
	message string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("My mistake is %v", ce.message)
}
func main() {
	c2 := customErr{message: "Don't Staying at Home"}
	foo(c2)
}
func foo(e error) {
	fmt.Println(e, "and", e.(customErr).message)
}
