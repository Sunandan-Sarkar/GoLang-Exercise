package main

import (
	"fmt"
)

func main() {
	cd := make(chan int)
	go func() {
		cd <- 55
		close(cd)
	}()
	i, ok := <-cd
	fmt.Println(i, ok)
	i, ok = <-cd
	fmt.Println(i, ok)
}
