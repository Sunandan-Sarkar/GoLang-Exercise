package main

import "fmt"

type person struct {
	First string
}
type human interface {
	speak()
}

func main() {
	p1 := person{
		First: "Sarkar"}
	p1.speak()
	saySomething(&p1)
}

func (p *person) speak() {
	fmt.Println("I speak Estonian Language")
}

func saySomething(h human) {
	h.speak()
}
