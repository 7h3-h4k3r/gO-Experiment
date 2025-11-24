package main 

import (
	"fmt"
)

type Address struct {
	city string 
	state string 
	pincode string
}

func (home Address) Location(){
	fmt.Printf("This is my current Location is : (%s,%s,[%s])",home.city,home.state,home.pincode)
}

func ( home User) Location(){
	fmt.Printf("Override the function has been shadow")
}
type User struct {
	name string 
	email string 
	password string
	Address //coposition
}

func main(){
	auth := User{name:"dharani",email:"tharan@gmail.com",password:"134@ery",Address:Address{city:"Dubai",state:"Lexican",pincode:"90FE3"}}
	auth.Location()
}