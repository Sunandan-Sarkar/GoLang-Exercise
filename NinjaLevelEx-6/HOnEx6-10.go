package main

import (
	"fmt"
)

func main() {
	d := roo()
	fmt.Println(d())
	fmt.Println(d())
	d1:=roo()
	fmt.Println(d1())
	fmt.Println(d1())
	fmt.Println(d1())
	fmt.Println(d1())

}
func roo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
