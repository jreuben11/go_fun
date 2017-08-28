package main

import "os"

func main() {
	// fail fast on unexpected errors
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
