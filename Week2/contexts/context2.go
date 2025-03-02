package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped:", ctx.Err())
			return
		default:
			fmt.Println("Worker running...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main2() {                                      //TO RUN THIS JUST CHANGE MAIN2() TO MAIN()
	ctx, cancel := context.WithCancel(context.Background()) // Create cancellable context

	go worker(ctx) // Start goroutine

	time.Sleep(3 * time.Second) // Let the worker run for 3 seconds
	cancel()                    // Manually stop the worker
	time.Sleep(1 * time.Second)  // Allow time for cleanup
	fmt.Println("Main function finished")
}
