package main

import (
	"fmt"
	"time"
)

func main() {

	// No of workers which use parallelly for the whole program
	NO_OF_WORKERS := 5

	// Integer array with values to wait for a worker to finish a sub process
	arr := [8]int{3, 2, 1, 8, 7, 6, 5, 1}

	// Jobs channel to send data to sub process from main process
	jobs_chan := make(chan int, len(arr))

	// Results channel to receive data from sub process from main process
	results_chan := make(chan string, len(arr))

	// Create workers in the concurrent go mode
	for w := 0; w < NO_OF_WORKERS; w++ {
		// Call function with Jobs channel and WaitGroup's Address
		go Process(jobs_chan, results_chan)
	}

	// Loop through the array
	for _, job := range arr {
		// Send a job data into the Jobs channel
		jobs_chan <- job
	}
	// Close the unused channel after sending all the data
	close(jobs_chan)

	// Loop in the number/times of arr
	for range arr {
		// Print a data from result channel
		fmt.Print(<-results_chan)
	}

}

// Function to process some instructions with Jobs channel and WaitGroup's Address
func Process(jobs_chan <-chan int, results_chan chan<- string) {

	// Loop and get all the data from the Job channel
	for job := range jobs_chan {

		// Sleep the process for job data second amount
		time.Sleep(time.Duration(int(time.Second) * job))

		// Send a result to the result channel
		results_chan <- fmt.Sprintf("Slept for %ds\n", job)

	}
}
