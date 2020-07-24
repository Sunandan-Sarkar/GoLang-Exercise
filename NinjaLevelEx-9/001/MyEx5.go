package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var np sync.WaitGroup

	fmt.Println("1st GR:", runtime.NumGoroutine())
	var increment int64
	ps := 100
	np.Add(ps)
	for i := 0; i <ps; i++ {
		go func() {
			atomic.AddInt64(&increment,1)
			fmt.Println(atomic.LoadInt64(&increment))
			np.Done()
		}()
	}

	fmt.Println("2nd GR:", runtime.NumGoroutine())
	np.Wait()
	fmt.Println("3rd GR:", runtime.NumGoroutine())
	fmt.Println("End value:", increment)
}


