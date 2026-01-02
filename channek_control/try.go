package main

func f() (result int) {
    defer func() {
        result++
    }()
    return 1
}

func main() {
    println(f())
}