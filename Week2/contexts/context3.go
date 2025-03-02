package main

import (
	"context"
	"fmt"
)

func processRequest(ctx context.Context) {
	userID := ctx.Value("userID") // Retrieve value from context
	fmt.Println("Processing request for user:", userID)
}

func main() {
	ctx := context.WithValue(context.Background(), "userID", 12345) // Store a value
	processRequest(ctx)
}
