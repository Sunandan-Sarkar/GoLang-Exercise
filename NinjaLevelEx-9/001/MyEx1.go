package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Before CPUs:", runtime.NumCPU())
	fmt.Println("Before Goroutines:", runtime.NumGoroutine())
	var np sync.WaitGroup
	np.Add(2)
	go func() {
		fmt.Println("I'm from thing One")
		np.Done()
	}()
	go func() {
		fmt.Println("I'm from thing Two")
		np.Done()
	}()
	fmt.Println("Mid CPUs:", runtime.NumCPU())
	fmt.Println("Mid Goroutines:", runtime.NumGoroutine())

	np.Wait()
	fmt.Println("About to exit")
	fmt.Println("After CPUs:", runtime.NumCPU())
	fmt.Println("After Goroutines:", runtime.NumGoroutine())
}
