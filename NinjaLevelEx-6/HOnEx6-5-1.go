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
type SHAPE interface {
	AREA() float64
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

	x1 := s1.AREA()
	fmt.Println(x1)
	x2 := c1.AREA()
	fmt.Println(x2)

	INFO(s1)
	INFO(c1)
}

func (s SQUARE) AREA() float64 {
	return s.length * s.wide
}
func (c CIRCLE) AREA() float64 {
	return math.Pi * (c.radious * c.radious)
}
func INFO(s SHAPE) {
	fmt.Println(s.AREA())
}
