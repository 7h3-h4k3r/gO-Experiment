package main 

import (
	"fmt"
)

type ERRFound struct{
	errString string 
}


func (e ERRFound) Error() string {
	return fmt.Sprintf("This is Error of the Current state",e.errString)
}


func validation(age int) error{
	if(18 > age){
		return ERRFound{errString:"Error on age"}
	}
	return nil 
}

func main(){

	err:= validation(10)
	if err!=nil{
		if value ,ok := err.(ERRFound);ok{
			fmt.Println("custom errr bro %v",value.Error())
		}
	
	}
}