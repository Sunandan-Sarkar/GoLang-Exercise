package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Outside a goroutine")
	go func() {
		fmt.Println("Inside a goroutine")
	}()
	//time.Sleep(time.Second)
	fmt.Println("Outside again goroutine")
	runtime.Gosched()
}
