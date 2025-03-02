package main

import (
	"context"
	"fmt"
	"time"
)

func task(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task stopped:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main4() {
	deadline := time.Now().Add(2 * time.Second) // Set deadline 2s from now
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go task(ctx)

	time.Sleep(3 * time.Second) // Let it run
	fmt.Println("Main function finished")
}
