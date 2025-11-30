package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()
	select {
	case <-ctx.Done():
		fmt.Println("Opration canceled :", ctx.Err())
	}
	fmt.Println("Hello, 世界")
}