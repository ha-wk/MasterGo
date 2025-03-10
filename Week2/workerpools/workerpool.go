package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function to process tasks
func worker2(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second) // Simulate work
		results <- job * 2      // Example: Multiply job by 2
	}
}

func main3() {
	const numWorkers = 3 // Number of workers
	const numJobs = 5    // Number of jobs

	jobs := make(chan int, numJobs)    // Channel for jobs
	results := make(chan int, numJobs) // Channel for results
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker2(i, jobs, results, &wg)
	}

	// Send jobs to workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // No more jobs to send

	// Wait for workers to finish
	wg.Wait()
	close(results) // Close results channel

	// Print results
	for result := range results {
		fmt.Println("Result:", result)
	}
}
