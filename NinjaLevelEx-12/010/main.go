package main

import (
	"fmt"
	"github.com/Sunandan-Sarkar/GoLang-Exercise/NinjaLevelEx-12/010/dog"
)

type robin struct {
	Name string
	age  int
}

func main() {
	robin := robin{
		Name: "Robin99",
		age:  dog.Years(10),
	}
	fmt.Println(robin)
}
