package main 

import (
	"fmt"
)

type call interface{
	Log()
}

type syslog struct{}

func (sys syslog) Log(){
	fmt.Printf("System has Been running on Port 80")
}

type computer struct{
	syslog
}

func getOpration(c call){
	c.Log()
}


func main(){
	process_id := computer{}
	getOpration(process_id)
}