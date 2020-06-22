package main

import (
	"fmt"
)

func main() {
	y := []string{"sun", "moon", "star", "rocket", "foot", "face"}
	fmt.Println(y)
	fmt.Println(cap(y))
	fmt.Println(len(y))
	for i,v:=range y{
		fmt.Println(i,v)
	}
	fmt.Println("Let's check the alternative method")
	//alternative method
	for i:=0;i<=5;i++{
		fmt.Println(i,y[i])
	}
	fmt.Println("How's it? Are you enjoying it?")
}
