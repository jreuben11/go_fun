package main

import "fmt"

func main() {
	// queue channel
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// range iterates over each element as it’s received from queue - channel close  --> iteration terminates
	for elem := range queue {
		fmt.Println(elem)
	}
}
