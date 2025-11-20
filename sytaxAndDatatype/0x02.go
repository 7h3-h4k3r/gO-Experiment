package main 

import (
	"fmt"
)
func example(){
	//redec , multi dec , shadow 
	var name int = 10
	if true{
		name := 20
		fmt.Println(name)
	}
	fmt.Println(name)
}
func main(){
	//Explicit Dec 
	var num  int = 34
	var value float64
	value = 90.00

	// inferred Dec

	c := 89.3
	s := "DharaiTharan"
	arr := []int{1,2,3,4}

	example()
	
}