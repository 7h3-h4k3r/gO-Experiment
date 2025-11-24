package main

import (
	"fmt"
)

type shape interface{
	area() interface{}
}

type circle struct{
	radies float64
}

func (c circle) area() interface{}{
	fmt.Printf("this is from cir function")
	return c.radies * c.radies
}

type rect struct{
	hight int 
	weight int 
}

func (r rect) area() interface{}{
	fmt.Printf("this from rectange function ")
	return r.hight * r.weight
}


func find_shape(s shape){
	fmt.Printf("Find Area From %v\n",s.area())
}

func main(){
	cir := circle{
		radies:194.5}

	rec := rect{hight:8,weight:90}

	find_shape(cir)
	find_shape(rec)
}


