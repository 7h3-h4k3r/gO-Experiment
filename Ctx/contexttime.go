package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer func() {
		cancel()
	}()
	select {
	case <-time.After(5 * time.Second):
		fmt.Printf("Time 5 sec has Been finished")
	case <-ctx.Done():
		fmt.Println("Opration canceled :", ctx.Err())
	}
	fmt.Println("Hello, 世界")
}
