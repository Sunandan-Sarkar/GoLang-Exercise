package main

import (
	"fmt"
)
func main(){
	c:=make(chan int)

	go func(){
		c<-99
		c<-100
	}()
	fmt.Println(<-c)
	fmt.Println(<-c)
}

