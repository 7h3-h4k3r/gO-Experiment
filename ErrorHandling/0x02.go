package main

import (
	"fmt"

)

func unsage(should bool) {

	defer func(){
		if r:= recover(); r != nil{
			fmt.Println("Recovered from panic :",r)
		}
	}()

	if should{
		panic(" i am stop from panic bro")
	}

	fmt.Print("but na enga erruka")
}

func main(){
	unsage(true)
}