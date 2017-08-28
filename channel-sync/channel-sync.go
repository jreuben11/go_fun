package main

import "fmt"
import "time"

// done channel will be used to notify another goroutine that this function’s work is done
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	// Send a value to notify done.
	done <- true
}

func main() {
	// Start a worker goroutine, giving it the channel to notify on
	done := make(chan bool, 1)
	go worker(done)

	// Block until receive a notification from the worker on the channel
	<-done
}
