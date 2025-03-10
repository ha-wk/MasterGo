//1️⃣ Worker Pool Implementation

//Write a worker pool where 4 workers process 10 numbers and return their squares.

//2️⃣ Fan-out Pattern

//Implement 4 goroutines that fetch data from a common channel and process it.


//These two questions can be demonstrated via single code base...here fan-out pattern can be seen by Channel jobs



package main

import (
	"fmt"
	"sync"
	"time"
)

func worker3(w int,jobs<-chan int,results chan<- int,wg *sync.WaitGroup){
    defer wg.Done()
	for j:=range jobs{
      fmt.Printf("Worker %d is processing job with number %d",w,j)
	  fmt.Println()
	  time.Sleep(time.Second)
	  results <- j*j
	}

}

func main(){
	jobs:=make(chan int,10)
	results:=make(chan int,10)
    var wg sync.WaitGroup

	for w:=1;w<=4;w++{
        wg.Add(1)
		go worker3(w,jobs,results,&wg)
    }

	// Send 10 jobs
	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	close(jobs) // Close jobs channel so workers know when to stop

	// Wait for all workers to complete
	wg.Wait()
	close(results) // Now safe to close results channel

	for r:= range results{
		fmt.Printf("Square of numbers are as follows %d",r)
		fmt.Println()
	}
}



//2️⃣ Fan-out Pattern

//Implement 4 goroutines that fetch data from a common channel and process it.

