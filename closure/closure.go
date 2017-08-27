package main

import "fmt"

// intSeq returns another function, which closes over the variable i
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	// call intSeq, assigning the function result to nextInt. This function value captures its own i value, which will be updated each time we call nextInt.
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// To confirm that the state is unique to that particular function, create and test a new one.
	newInts := intSeq()
	fmt.Println(newInts())
}
