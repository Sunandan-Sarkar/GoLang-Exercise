package main

import (
	"fmt"
)
type person struct {
	Saying string

}
type human interface {
	speak()
}

func main(){
p1:=person{
	Saying: "I speak Estonian"}

p1.speak()
saySomething(&p1)
}

func(p *person) speak(){
fmt.Println(*p)
}

func saySomething(h human){
h.speak()
}
