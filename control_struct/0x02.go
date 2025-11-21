package main

import (
	"fmt"
)

func main(){
	// C - style 
	for i := 0;i<5;i++{
		fmt.Print(i," ->");
	}
	fmt.Println();

	var num int = 0;
	//while 
	for num < 5{
		fmt.Print(num);
		num++
	}
	fmt.Println()
	//Range over slice
	nums := []int{10,30,30,46}
	for index, value := range nums {
		fmt.Println(index,value)
	}


	// same use case of Break and Continue

	
}
