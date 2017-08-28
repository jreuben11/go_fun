package main

import "fmt"

// int parameter - arguments passed by value. (copy)
func zeroval(ival int) {
	ival = 0
}

// *int parameter - *iptr in function body dereferences
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// &i syntax - pass the memory address
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// print Pointer address
	fmt.Println("pointer:", &i)
}
