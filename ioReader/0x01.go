package main 

import (
	"fmt"
	"io"
)

type simpleSlicedata struct{
	data string 
	pos int 
} 

func (r *simpleSlicedata) testing(p []byte) (int, error) {
	if r.pos >= len(r.data){
		return 0,io.EOF 
	}

	count := copy(p,r.data[r.pos:])
	r.pos+=count 
	return count,nil
}

func main(){
	src := &simpleSlicedata{data:"Hello ! form new Galaxy for Go language"}
	buf := make([]byte,8)
	for {
		n , err := src.testing(buf)
		if err == io.EOF{
			break
		}
		fmt.Printf("Read %d Bytes: %s\n",n,string(buf[:n]))
	}
}