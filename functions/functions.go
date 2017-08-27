package main

import "fmt"

// a function that takes two ints and returns an int.
func plus(a int, b int) int {
	// Go requires explicit returns
	return a + b
}

// With multiple consecutive parameters of same type, can omit type name up to final parameter
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	// Call a function
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
