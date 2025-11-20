package main

import (

	"fmt"
)

func main(){
	f:= float64(10)
	it := int(f)

 	//type asseration 
	var value intreface{} = "hello"
	s , ok := value.(string)
}