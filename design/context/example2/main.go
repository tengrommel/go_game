package main

import (
	"context"
	"time"
	"fmt"
)

// context的使用规范

func main() {
	// Create a context that is cancelled only manually.
	ctx, cancel := context.WithCancel(context.Background())

	// The cancel function must be called regardless of the outcome.
	defer cancel()

	// Ask the goroutine to do some work for us.
	go func() {
		for{
			// Simulate work.
			time.Sleep(50 * time.Millisecond)
			// Report the work is done.
			cancel()
		}
	}()

	// Wait for the work to finish. If it takes too long move on.
	select {
	case <- time.After(100 * time.Millisecond):
		fmt.Println("moving on")
	case <- ctx.Done():
		fmt.Println("work complete")
	}

}
