package main

import "fmt"

func main() {
	cm := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			cm <- i
		}
		close(cm)
	}()
	for v := range cm {
		fmt.Println(v)
	}
	fmt.Println("About to exit")
}
