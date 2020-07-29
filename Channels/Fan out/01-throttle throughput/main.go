package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c3 := make(chan int)
	c4 := make(chan int)

	go send(c3)

	go fanout(c3, c4)

	for v := range c4 {
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
func fanout(c3, c4 chan int) {
	var wg sync.WaitGroup
	const gs = 10
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			for v := range c3 {
				func(v2 int) {
					c4 <- timeconsumingwork(v2)
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c4)

}

func timeconsumingwork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
