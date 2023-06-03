package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// No of workers which use parallelly for the whole program
	NO_OF_WORKERS := 5

	// Integer array with values to wait for a worker to finish a sub process
	// arr := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	arr := [8]int{3, 2, 1, 8, 7, 6, 5, 1}

	// Waitgroup to wait till the sub process finish not to end the entire program
	var wg sync.WaitGroup

	// No of WaitGroup counter to wait
	wg.Add(len(arr))

	// Wait at the end of the program
	defer wg.Wait()

	// Jobs channel to send data to sub process from main process
	jobs_chan := make(chan int, len(arr))

	// Create workers in the concurrent go mode
	for w := 0; w < NO_OF_WORKERS; w++ {
		// Call function with Jobs channel and WaitGroup's Address
		go Process(jobs_chan, &wg)
	}

	// Loop through the array
	for _, job := range arr {
		// Send a job data into the Jobs channel
		jobs_chan <- job
	}

	// Close the unused channel after sending all the data
	close(jobs_chan)

}

// Function to process some instructions with Jobs channel and WaitGroup's Address
func Process(jobs_chan <-chan int, wg *sync.WaitGroup) {

	// Loop and get all the data from the Job channel
	for job := range jobs_chan {

		// Sleep the process for job data second amount
		time.Sleep(time.Duration(int(time.Second) * job))

		// Print which Job has slept
		fmt.Printf("Slept for %ds\n", job)

		// Call WaitGroup Done for a WaitGroup counter
		wg.Done()
	}
}
