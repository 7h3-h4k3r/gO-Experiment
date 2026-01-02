package main

import (
    "fmt"
    "example.com/shared-lib"
)

func main() {
    fmt.Println(sharedlib.Greet("AuthService"))
}
