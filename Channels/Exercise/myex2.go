package main

import (
	"fmt"
)
func main(){
	c:=make(chan int,2)
	c<-88
	c<-77
	fmt.Println(<-c)
	fmt.Println(<-c)
}

