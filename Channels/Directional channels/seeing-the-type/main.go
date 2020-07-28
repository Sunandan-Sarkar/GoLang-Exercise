package main

import (
	"fmt"

)
func main(){
	c:=make(chan int,2)
	c<-55
	c<-66
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Printf("%T\t",c)
}

