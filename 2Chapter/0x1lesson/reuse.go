package main

import "fmt"

type Logger struct {
	Level string
}
type service struct {
	Logger
	name string
}

func (l Logger) log(message string) {
	fmt.Printf("[%s] %s\n", l.Level, message)

}

func main() {
	srv := service{Logger: Logger{Level: "DUBUG"}, name: "auth"}
	srv.log("running")
}
