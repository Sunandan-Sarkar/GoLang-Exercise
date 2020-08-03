package main

import (
	"fmt"
)

type movie struct {
	name    string
	release int
}

func main() {
	dark := movie{
		name:    "Dark",
		release: 2017,
	}
	fmt.Println(dark)
}
