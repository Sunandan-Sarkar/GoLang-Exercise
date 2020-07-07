package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Dr.",
		last:  "Strange",
		age:   46,
	}
	//p2:=person{
	//	first: "Usain",
	//	last:  "Bolt",
	//	age:   38,
	//}
	fmt.Println(p1)
	//fmt.Println(p2)
	p1.speak()
}
func (s person) speak() {
	fmt.Println(s.first, s.last, s.age)
}
