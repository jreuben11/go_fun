package main

import "fmt"

// use jobs channel to communicate work to be done from the main() goroutine to a worker goroutine.
// When no more jobs for the worker, close the jobs channel.
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// worker goroutine -repeatedly receives from jobs with j, more := <-jobs. more value will be false if jobs has been closed and all values in the channel have already been received.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// send 3 jobs to worker over jobs channel, then close it
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// await  worker using synchronization approach
	<-done
}
