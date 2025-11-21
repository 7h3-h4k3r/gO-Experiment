package main

import (
	"fmt"
)

func main(){
	value := 3

	switch value{
		case 1:
			fmt.Println("one_value")
		case 2,3,4:
			fmt.Println("is in list")
		default:
			fmt.Println("no in that")
	}

	var i interface{} = 90.34
	switch i.(type){

		case float64:
			fmt.Println("this is float")
		case string:
			fmt.Println("string value broo")
		case int:
			fmt.Println(" yes its a INterfer")
		default:
			fmt.Println("not type its like union code")
	}
}