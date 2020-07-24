package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutine:", runtime.NumGoroutine())

	var increment int64
	const gp = 100
	var wp sync.WaitGroup
	wp.Add(gp)

	for i := 0; i < gp; i++ {
		go func() {
			atomic.AddInt64(&increment, 1)
			fmt.Println("Increments\t", atomic.LoadInt64(&increment))
			runtime.Gosched()
			wp.Done()
		}()
	}
	fmt.Println("Sum:", increment)
	wp.Wait()
}
