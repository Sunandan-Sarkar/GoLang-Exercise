package main

import "fmt"

func main() {
	fmt.Println("This is func expression")
	f := func(x int) {
		fmt.Println("Sundan was born in", x)
	}
	f(1987)
	f2 := func(y string) {
		fmt.Println("Sundan was bon in", y)
	}
	f2("Bangladesh")
	f3 := func(z []string) {
		fmt.Println("Sundan was", z)
	}
	f3([]string{"Awesome", "Honest", "Kind"})
}
