package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)

	go receive(eve, odd, quit)

	fmt.Println("About to exit")

}
func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the eve channel", v)
		case v := <-o:
			fmt.Println("From the odd channel", v)
		case <-q:
			return
		}

	}
}
