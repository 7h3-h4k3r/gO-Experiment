package main

import (
    "fmt"
)

var count int

func increment() {
    count = count + 1
}

func main() {
    for i := 0; i < 5; i++ {
        go increment()
    }

    fmt.Println(count)
}
