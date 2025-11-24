package main 

import (
	"fmt"
)


type wheel interface{
	move()
}

type car struct{
	model string
}

func (c *car ) move(){
	fmt.Printf("[%s] START Engine Ready to go The Race yruumm yruummm\n",c.model)
}

type bike struct{
	model string
}

func (b bike) move(){
	fmt.Printf("[%s] START Engine Ready to fo the Track Race All Good\n",b.model)
}


func racer(w wheel){
	w.move()
}


func main(){
	var v wheel
	bi := bike{model:"Ninja H2 ^6"}
	v = bi 
	v.move() 
	

	ca := car{model:"BMW s series 45023"}
	v = &ca
	v.move()
	
}
