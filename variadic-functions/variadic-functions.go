package main

import "fmt"

// take an arbitrary number of ints as arguments.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	// calling Variadic functions
	sum(1, 2)
	sum(1, 2, 3)

	// apply a slice to a variadic function
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
