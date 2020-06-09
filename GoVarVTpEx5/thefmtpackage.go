package main

import (
	"fmt"
)
var y =42
func main (){
	fmt.Println(y)
fmt.Printf("%T\n",y)
	fmt.Printf("%#v\n",y)
	fmt.Printf("%x\n",y)
	fmt.Printf("%d\n",y)
	fmt.Printf("%o\n",y)
	fmt.Printf("%U\n",y)
	y=911
	fmt.Printf("%#v\t%x\t%b\t",y,y,y)
}

