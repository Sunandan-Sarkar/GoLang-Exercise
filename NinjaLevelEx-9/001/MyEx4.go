package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var np sync.WaitGroup

	fmt.Println("1st GR:", runtime.NumGoroutine())
	increment := 0
	ps := 100
	np.Add(ps)
	var mt sync.Mutex
	for i := 0; i <ps; i++ {
		go func() {
			mt.Lock()
			v := increment
			//runtime.Gosched()
			v++
			increment = v
			fmt.Println(increment)
			mt.Unlock()
			np.Done()
		}()
	}

	fmt.Println("2nd GR:", runtime.NumGoroutine())
	np.Wait()
	fmt.Println("3rd GR:", runtime.NumGoroutine())
	fmt.Println("End value:", increment)
}

