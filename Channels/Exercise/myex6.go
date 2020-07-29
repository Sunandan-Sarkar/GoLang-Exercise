package main

import "fmt"

func main(){
	d:=make(chan int)
	q:=make(chan int)
	go func() {
		for i:=0;i<10;i++{
			fmt.Println(<-d)
		}
		q<-0
	}()
	boo(d,q)
}

func boo(d,q chan int){
	x:=3
	for{
		select {
		case d <-x:
			x+=x
		case <-q:
			fmt.Println("About to exit")
			return
		}
	}
}
