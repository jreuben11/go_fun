package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// state will be owned by a single goroutine to guarantee data is is thread-safe
// to read / write state, other goroutines will send messages to owning goroutine & receive corresponding replies.

// readOp / writeOp structs encapsulate requests + channel for owning goroutine to respond.
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	// count operations
	var readOps uint64 = 0
	var writeOps uint64 = 0

	// reads / writes channels - used by other goroutines to issue read and write requests
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// goroutine that owns the state (a map). repeatedly selects on reads & writes channels, responding to requests as they arrive.
	// A response is executed by first performing the requested operation and then sending a value on the response channel resp
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// start 100 goroutines to issue reads to the state-owning goroutine via the reads channel.
	// Each read constructs a readOp, sends it over the reads channel, and wait to receive the result over the provided resp channel.
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// start 10 writes using a similar approach.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Let the goroutines work for a second.
	time.Sleep(time.Second)

	// Finally, capture and report the op counts.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
