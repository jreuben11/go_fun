package main

import "fmt"

func main() {
	// type of elements and length are both part of the array’s type. By default an array is zero-valued
	var a [5]int
	fmt.Println("emp:", a)

	// set / get a value at an index
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// builtin len
	fmt.Println("len:", len(a))

	// syntax to declare and initialize an array in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// can compose types to build multi-dimensional data structures.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
