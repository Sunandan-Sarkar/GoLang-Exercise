package main

import (
	"fmt"
)

func main() {
	tal("Sunandan", "Sarkar", "Novi", "Sarkar")
	boo("Jesmine", 4, 6, 8, 10, 12)
	//final argument(l ...string) or (bm ...int) is assignable to a slice type []T
	Greeting("nobody")
	//1st call,within Greeting, who will have the value nil in the first call
	Greeting("hello:", "Joe", "Anna", "Eileen")
	// 2nd call and []string{"Joe", "Anna", "Eileen"} in the second.

	s := []string{"James", "Jasmine", "Moon"}
	Greeting("goodbye:", s...)
	//3rd call, within Greeting, who will have the same value as s with the same underlying array. because both followed by ...
}
func tal(f string, l ...string) {
	fmt.Println(f, l)
	fmt.Printf("%T\n%T\n", f, l)
}
func boo(bn string, bm ...int) {
	fmt.Println(bn, bm)
	fmt.Printf("%T\n%T\n", bn, bm)
}
func Greeting(prefix string, who ...string) {
	fmt.Println(prefix, who)
	fmt.Printf("%T\n%T\n", prefix, who)
}
