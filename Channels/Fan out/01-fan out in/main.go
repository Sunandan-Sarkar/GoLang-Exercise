package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go send(c1)

	go fanoutin(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("Done!")
}

func send(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
func fanoutin(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		go func(v2 int) {
			wg.Add(1)
			c2 <- timeconsumingwork(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}
func timeconsumingwork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
