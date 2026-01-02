package main 

import (
	"fmt"
	"time"
)

func worker(w int,jobs <-chan int , results chan<- int){
	for valu := range jobs{
		fmt.Printf("Working thread %d start job is %d\n",w,valu)
		time.Sleep(time.Second)
		fmt.Printf("Thread Finised %d stop jon is %s\n",w,valu)
		results <-valu * 2 
	} 
}

func main(){
	const variable   = 5 
	jobs := make(chan int ,variable)
	results := make(chan int,variable)
	for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send jobs and close channel to signal no more jobs
    for j := 1; j <= variable; j++ {
        jobs <- j
    }
    close(jobs)

    // Collect results
    for a := 1; a <= variable; a++ {
        res := <-results
        fmt.Printf("Result received: %d\n", res)
    }
}