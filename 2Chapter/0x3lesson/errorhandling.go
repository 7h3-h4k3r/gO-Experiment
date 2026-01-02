package main

import (
	"fmt"
)

type tempFile struct{
	Message string
}

func (t tempFile) Temporary()bool{
	return true
}

func (t tempFile) Error() string {
	return t.Message
}
func shouldDo(value bool) error{
	if value{
		return tempFile{Message:"Haveing to Use a some modularity Err"}
	}
	return nil 
}

func main(){
	err := shouldDo(true)
    if err != nil {
        fmt.Println("Error:", err)
        type temporary interface {
            Temporary() bool
        }
        if te, ok := err.(temporary); ok && te.Temporary() {
            fmt.Println("This is a temporary error; can retry operation.")
        }
    }
}