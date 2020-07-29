package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("Context:\t", ctx)
	fmt.Println("Context Error:\t", ctx.Err())
	fmt.Printf("Context type:\t %T\n",ctx)
	fmt.Println("--------------------------")

	ctx,_=context.WithCancel(ctx)

	fmt.Println("Context:\t", ctx)
	fmt.Println("Context Error:\t", ctx.Err())
	fmt.Printf("Context type:\t %T\n",ctx)
	fmt.Println("--------------------------")
}
