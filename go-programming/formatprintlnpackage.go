package main

import "fmt"

func main(){
	//n, err:=fmt.Println("Hello, playground")
	//fmt.Println(n)
	//fmt.Println(err)
	// It returns the number of bytes written and any write error encountered
	n, _:=fmt.Println("Hello, playground")
	fmt.Println(n)
//_ throw away the return for error
}
