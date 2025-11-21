package main

import (
	"fmt"
)

func add(a int ,b int) int{
	return a + b
}
// muliple return value 

func swap(x  ,y int ) (int,int){
	return y,x 
}

//name return values

func getData(age int)(name string ,inc int){
	
	name = "yes auth granted"
	
	inc = age + 3
	return 
}

// variafic functino

func sum(nums ...int) int{
	var total int = 0;
	for _,n := range nums{
		total+=n 
	}
	return total
}


func anaonymous_functino(fn func(int , int ) int  ,x ,y int ) int {	
	return fn(x,y)
}

func test_clouser() func() int{

	x:=0 
	return func() int {
		x++
		return x
	}
}

func factorial(n int) int {
	if n <=1{
		return 1 
	}

	return n * factorial(n-1)
}
func main(){
	fmt.Printf("3 + 5 = %d ",add(3,5))
	intervalue , string_value := swap(3,7)
	fmt.Print(intervalue,string_value)
	name , vailed_age  := getData(18)
	fmt.Printf("validate age is [%d] %v\n",vailed_age,name)
	fmt.Print(sum(1,2,3,4,5));
	fmt.Println()
	fmt.Print(anaonymous_functino(add,108,5))
	c:= test_clouser()
	fmt.Println()
	fmt.Print(c())
	fmt.Print(c())
	fmt.Println()
	fmt.Print(factorial(5))
}
