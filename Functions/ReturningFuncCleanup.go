package main

import "fmt"

func main() {
	x := bar()
	fmt.Printf("%T\n", x)

	// x()Run func
	fmt.Println(x())

	//More cleanup
	fmt.Println(goo()())
}

func bar() func() int {
	return func() int {
		return 451
	}
}

//More cleanup
func goo() func() string {
	return func() string {
		return "Limbaroyati"
	}
}
