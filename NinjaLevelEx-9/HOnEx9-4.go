package main

import (
	"fmt"
	"runtime"
	"sync"
)
func main(){
	fmt.Println("CPUs:",runtime.NumCPU())
	fmt.Println("Goroutine:",runtime.NumGoroutine())

	increment:=0
	const gp=100
	var wp sync.WaitGroup
	wp.Add(gp)

	var pu sync.Mutex

	for i:=0; i<gp; i++{
		go func(){
			pu.Lock()
			v:=increment
			v++
			increment=v
			fmt.Println(increment)
			pu.Unlock()
			wp.Done()
		}()
	}
	fmt.Println("Sum:",increment)
	wp.Wait()
}

