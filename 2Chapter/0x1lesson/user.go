package main
 
import (
	"fmt"
)


type User struct{
	name string
	id int 
	active bool
}

func (u User) is_active() bool{
	return u.active
} 

func (u *User) deactive_at(){
	u.active = false
}

func main(){
	u:= User{
		id:1,
		name:"dharani",
		active:true}
	fmt.Println("Data ",u)
	fmt.Println("Before deactive ",u.is_active())
	u.deactive_at()
	fmt.Println("after deactive ",u.is_active())
}