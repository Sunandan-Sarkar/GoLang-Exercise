package main

import "fmt"

func main(){
	ct:=make(chan<- int)
	go func() {
		ct<-11
	}()
	fmt.Println(<-ct)

	fmt.Println("--------")
	fmt.Printf("%T\t",ct)
	//error: “invalid operation: <-cs (receive from send-only type chan<- int)”
}
