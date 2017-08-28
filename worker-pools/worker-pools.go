package main

import "fmt"
import "time"

// workers - will receive work on jobschannel & send results on results. sleep a second per job to simulate
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	//send workers work and collect their results - make 2 channels for this.
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// start up 3 workers, initially blocked because  no jobs yet.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// send 5 jobs, then close channel
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// collect all results of the work
	for a := 1; a <= 5; a++ {
		<-results
	}
}
