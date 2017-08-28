// When using channels as function parameters, can specify if a channel is meant to only send or receive values

package main

import "fmt"

// function only accepts a channel for sending values
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//function accepts one channel for receives and a second for sends
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
