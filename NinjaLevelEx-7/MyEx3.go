package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println("X befor",&x)
	fmt.Println("X befor",x)
	foo(&x)
	fmt.Println("X after",&x)
	fmt.Println("X after",x)
}
func foo(y *int) {
	fmt.Println("Y befor",y)
	fmt.Println("Y befor",*y)
	*y=99
	fmt.Println("Y after",y)
	fmt.Println("Y after",*y)
}
