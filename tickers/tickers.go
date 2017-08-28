package main

import "time"
import "fmt"

func main() {
	// do something repeatedly at regular intervals
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		// use the range builtin on the channel to iterate over the values as they arrive
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// stop Ticker
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
