package main

import (
	"errors"
	"fmt"
	"os"
)

func Readfile(file_name string)([]byte ,error){
	data , err := os.ReadFile(file_name)
	if err != nil{
		return nil,err
	}
	return data,nil 
}


func readfile(file_name string) error{
	_ ,err  := os.ReadFile(file_name)
	if err!=nil{
		return fmt.Errorf("load %s :%w ",file_name,err)
	}
	return nil
}

func main(){
	// data , err := Readfile("config.json")
	// if err != nil{
	// 	fmt.Println("config file not int that place",err)
	// 	os.Exit(-1)
	// }
	// fmt.Printf(string(data))

	err := readfile("new.json")
	if err !=nil{
		var pathErr *os.PathError 
		if errors.As(err,&pathErr){
			fmt.Println("Path error on ",pathErr.Path)
		}
		fmt.Println("Error :",err);
	}
}	