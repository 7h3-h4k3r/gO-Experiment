package main

import (
    "fmt"
    "time"
)

func main() {
    counter := 0
    for i := 0; i < 5; i++ {
        go func() {
            counter++
        }()
    }
    time.Sleep(time.Second)
    fmt.Println("Counter:", counter)
}