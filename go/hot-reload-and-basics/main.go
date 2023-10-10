package main

import (
	"fmt"
	"sync"
)

func main() {
	// Start the terminal app in a Goroutine
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		console()
	}()
	// Wait for the terminal app to complete
	wg.Wait()
	fmt.Println("Exiting the main function.")
}
