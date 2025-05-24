package main

// WaitGroup is used to wait for a group of goroutines to finish.

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal completion when done

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Duration(id) * 100 * time.Millisecond)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Launch 5 workers
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment counter before starting goroutine
		go worker(i, &wg)
	}

	fmt.Println("Waiting for all workers to finish...")
	wg.Wait() // Wait for all workers to complete
	fmt.Println("All workers completed!")
}
