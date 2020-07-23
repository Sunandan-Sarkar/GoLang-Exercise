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

	for i:=0; i<gp; i++{
		go func(){
			v:=increment
			runtime.Gosched()
			v++
			increment=v
			fmt.Println(increment)
			wp.Done()
		}()
	}
	fmt.Println("Sum:",increment)
	wp.Wait()
}

