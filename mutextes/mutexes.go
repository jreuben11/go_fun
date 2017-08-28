package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	// state - a map
	var state = make(map[int]int)

	// mutex to synchronize access to state
	var mutex = &sync.Mutex{}

	// track how many read / write ops
	var readOps uint64 = 0
	var writeOps uint64 = 0

	// start 100 goroutines to execute repeated reads against the state, once per millisecond in each goroutine.
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {

				// For each read, pick a key to access, Lock() the mutex ,read the value, Unlock() the mutex, & increment the readOps count.
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				// Wait a bit between reads.
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// start 10 goroutines to simulate writes
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	// wait a bit for work
	time.Sleep(time.Second)

	// Take and report final operation counts.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// With a final lock of state, show how it ended up.
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
