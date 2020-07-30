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
	go fanout(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("Done!")

}
func send(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)

}
func fanout(c1, c2 chan int) {
	var wg sync.WaitGroup
	const gs = 10
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			for v := range c1 {
				func(v2 int) {
					c2 <- timeconsumingwork(v2)
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)
}

func timeconsumingwork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
