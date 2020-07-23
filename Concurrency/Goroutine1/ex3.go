package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Number of Goroutine 1st:", runtime.NumGoroutine())

	wg.Add(5)
	go moo("Hello")

	//wg.Add(1)
	go moo("Football")
	//runtime.Gosched()
	go coo()

	fmt.Println("Number of Goroutine 2nd:", runtime.NumGoroutine())

	go func(p string) {
		fmt.Println(p)
		wg.Done()
	}("Goroutine Inside first")
	//runtime.Gosched()

	fmt.Println("Number of Goroutine 3rd:", runtime.NumGoroutine())

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("Goroutine Inside second:", i)
		}
		wg.Done()
	}()
	//runtime.Gosched()

	fmt.Println("Number of Goroutine 4th:", runtime.NumGoroutine())
	wg.Wait()
	//time.Sleep(time.Second)

}
func moo(f string) {
	for i := 0; i < 5; i++ {
		fmt.Println(f, ":", i)
	}
	wg.Done()
}
func coo() {
	for i := 0; i < 4; i++ {
		fmt.Println("COO:", i)
	}
	wg.Done()
}
