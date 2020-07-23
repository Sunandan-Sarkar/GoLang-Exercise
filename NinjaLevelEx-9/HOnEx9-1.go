package main

import (
	"fmt"
	"runtime"
	"sync"
)
func main(){
fmt.Println("CPUs",runtime.NumCPU())
fmt.Println("Goroutine",runtime.NumGoroutine())

var mp sync.WaitGroup
	mp.Add(2)
go func(){
	for i:=0;i<5;i++{
		fmt.Println("My Score:",i)
	}
	mp.Done()
}()
go func(){
	fmt.Println("I goal one")
	mp.Done()
}()
	fmt.Println("Goroutine",runtime.NumGoroutine())
	mp.Wait()
}

