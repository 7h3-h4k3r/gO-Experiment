package main

import (
	"fmt"
)


func subtask(ch chan string){
	ch <-"epa epa than pa finished achi"
}


func subtas(ch chan string){
	subtask(ch)
}

func task(ch chan string){
	subtas(ch)
}

func main(){
	ch := make(chan string)

	go task(ch)

	msg := <-ch 
	fmt.Printf(msg)
}