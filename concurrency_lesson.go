// Example of using Goroutines and WaitGroups to manage concurrency

package main

import (
	"fmt"
	"sync"
	"time"
)

// simulateWork simulates a task that takes time (like an API call)
func simulateWork(id int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done() // Notify the WaitGroup when this function finishes

	fmt.Printf("Worker %d: Starting...\n", id)
	time.Sleep(time.Second * 2) // Sleep for 2 seconds
	
	result := fmt.Sprintf("Worker %d: Task Complete", id)
	results <- result // Send the result into the channel
}

func main() {
	var wg sync.WaitGroup
	results := make(chan string, 3) // A buffered channel to hold 3 results

	start := time.Now()

	// Launch 3 goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Tell the WaitGroup we are starting 1 task
		go simulateWork(i, &wg, results)
	}

	// Wait for all goroutines to finish in the background
	go func() {
		wg.Wait()
		close(results) // Close the channel when all workers are done
	}()

	// Collect results from the channel
	fmt.Println("Main: Waiting for workers...")
	for res := range results {
		fmt.Println(res)
	}

	fmt.Printf("Total time taken: %v\n", time.Since(start))
	fmt.Println("If this were sequential, it would take 6 seconds. With Goroutines, it takes ~2!")
}
