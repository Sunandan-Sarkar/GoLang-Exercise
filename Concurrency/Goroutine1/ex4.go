package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go loo()
	number := 0

	go func() {
		for i := 0; i <= 100; i++ {
			fmt.Println("Score Goal number:", i)
		}
		wg.Done()
	}()
	fmt.Println("My Score:", number)
	wg.Wait()

}
func loo() {
	for i := 0; i < 3; i++ {
		fmt.Println("Goroutine:", i)
	}
	wg.Done()

}
