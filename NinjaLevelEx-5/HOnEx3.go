package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	v1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "Red"},
		fourWheel: true,
	}
	v2 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Yellow"},
		luxury: false,
	}
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v1.color, v2.color, v1.fourWheel, v2.doors, v2.luxury, v1.doors)
}
