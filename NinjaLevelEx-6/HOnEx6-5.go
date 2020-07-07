package main

import (
	"fmt"
	"math"
)

type SQUARE struct {
	length float64
	wide   float64
}
type CIRCLE struct {
	radious float64
}

func main() {
	s1 := SQUARE{
		length: 34,
		wide:   75,
	}
	c1 := CIRCLE{
		radious: 19.34,
	}
	fmt.Println(s1, c1)
	s1.AREA()
	c1.AREA()
}

func (s SQUARE) AREA() {
	fmt.Println("Total area size ", s)
	fmt.Println("square area", s.length*s.wide, "sq.m")
}
func (c CIRCLE) AREA() {
	fmt.Println("Total size ", c)
	fmt.Println("circle area", math.Pi*(c.radious*c.radious))
}
