package main

import (
	"fmt"
)

func main() {
	//if x:= 42;x==2 { Didn't run because x not the same
	//	fmt.Println("007")
	//}
	if x := 42; x == 42 {
		fmt.Println("009")
	}

}
