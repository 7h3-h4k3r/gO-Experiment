package main

import 
(
	"fmt"
	"time"
)

func work(n int){
	fmt.Printf("start Working the %d\n",n);
	time.Sleep(time.Second)
	fmt.Printf("work done %d\n",n);
}

func main(){
	for i:=1;i<=5;i++{
		go work(i)
	}
	time.Sleep(2*time.Second)
	fmt.Printf("All is finished");
}