package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// call function in usual synchronous way
	f("direct")

	// invoke function in a goroutine, -  execute concurrently with the calling one
	go f("goroutine")

	// start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Scanln - read input
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
