package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "Iron",
		last:  "Man",
	}
	fmt.Println(p1)
	foo(&p1)

}
func foo(p *person) {
	p.first = "Tommy"
	p.last = "Green"
	fmt.Println(*p)
}
