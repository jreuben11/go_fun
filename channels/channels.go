package main

import "fmt"

func main() {
	// Create a new channel with make(chan val-type). Channels are typed
	messages := make(chan string)

	// Send a value into a channel using channel <-   eg from a new goroutine.
	go func() { messages <- "ping" }()

	// <-channel receives a value from the channel
	msg := <-messages
	fmt.Println(msg)
}
