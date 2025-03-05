package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // If context is canceled
			fmt.Println("Task stopped:", ctx.Err())
			return
		default:
			fmt.Println("Processing...")
			time.Sleep(1 * time.Second) // Simulating work
		}
	}
}

func main1() {
	// Create a context with timeout of 3 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // Ensure cancel is called to free resources

	go longRunningTask(ctx) // Start goroutine

	time.Sleep(5 * time.Second) // Let it run for 5 seconds (it will stop after 3s)
	fmt.Println("Main function finished")
}
