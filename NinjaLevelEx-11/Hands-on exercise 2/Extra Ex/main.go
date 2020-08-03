package main

import (
	"errors"
	"fmt"
)

var err = errors.New("Custom erros-My error")

func main() {
	fmt.Printf("%T\n", err)
	if err != nil {
		fmt.Println(err)
	}
}
