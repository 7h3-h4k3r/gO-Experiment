/* install go lib for bcrypt

install : bcrypt
go get golang.org/x/crypto/bcrypt

*/
package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func getHash(pass string) ([]byte , error){
	return bcrypt.GenerateFromPassword([]byte(pass),bcrypt.DefaultCost)
}

func verify_pass(hashpass []byte, pass string ) error{
	return bcrypt.CompareHashAndPassword(hashpass,[]byte(pass))
}

func main(){

	pass := "passw@6d"
	Hashed , ok := getHash(pass)
	if ok != nil{
		fmt.Println("getHash Error ::: ")
	}
	fmt.Println("Current Password ",pass)
	fmt.Println("pass hashword ",string(Hashed))
	if err := verify_pass(Hashed,pass); err == nil{
		fmt.Println("password Correct")
	}else{
		fmt.Println("password incorrect")
	}
}