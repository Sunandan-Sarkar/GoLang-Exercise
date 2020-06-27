package main

import (
	"fmt"
)

func main() {
	tal()
	tartu("I'm Sunandan, Nice to meet you!")
	s1 := saa("What's up!")
	fmt.Println(s1)
	n, m := vali("Sunandan ", " Sarkar ")
	fmt.Println(n, m)
}
func tal() {
	fmt.Println("Hello from Tallinn")
}
func tartu(t string) {
	fmt.Println("Hello,", t)
}
func saa(s string) string {
	return fmt.Sprint("Awesome!!! ", s)

}
func vali(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, "says, Hello")
	b := false
	return a, b
}
