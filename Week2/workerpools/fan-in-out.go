package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Worker function (Fan-out)
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500))) // Simulate work
		results <- job * 10 // Fan-in: Collecting results
	}
}

func main2() {
	rand.Seed(time.Now().UnixNano())

	jobs := make(chan int, 5)   // Job queue
	results := make(chan int, 5) // Merging results

	// Start multiple workers (Fan-out)
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // No more jobs

	// Receive results (Fan-in)
	for i := 1; i <= 5; i++ {
		fmt.Println("Final Result:", <-results)
	}
}
