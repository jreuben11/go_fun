package main

import "fmt"
import "time"
import "sync/atomic"

func main() {
	// use an unsigned integer to represent counter
	var ops uint64 = 0

	// simulate concurrent updates, with 50 goroutines that each increment the counter about once a millisecond.
	for i := 0; i < 50; i++ {
		go func() {
			for {

				// increment the counter - use AddUint64 against memory address of ops counter
				atomic.AddUint64(&ops, 1)

				// Wait a bit between increments.
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Wait a second to allow some ops to accumulate.
	time.Sleep(time.Second)

	//to safely use the counter while it’s still being updated by other goroutines, extract a copy of the current value via LoadUint64
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
