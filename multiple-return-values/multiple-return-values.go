package main

import "fmt"

// function signature - returns 2 ints.
func vals() (int, int) {
	return 3, 7
}

func main() {
	// assign the 2 different return values from the call with multiple assignment.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// for a subset of the returned values, use  blank identifier _.
	_, c := vals()
	fmt.Println(c)
}
