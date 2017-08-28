package main

import "fmt"
import "sort"

func main() {
	// Sort methods are specific to the builtin type and in-place

	// sorting strings
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// sorting ints.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// check if a slice is already in sorted order.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
