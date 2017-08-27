package main

import "fmt"

func main() {
	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
	// To create an empty slice with non-zero length, use  make.
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// set and get
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	// len
	fmt.Println("len:", len(s))

	// append
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	//  “slice” operator
	l := s[2:5]
	fmt.Println("sl1:", l)
	//  up to (but excluding)
	l = s[:5]
	fmt.Println("sl2:", l)
	// up from (and including)
	l = s[2:]
	fmt.Println("sl3:", l)

	// declare and initialize slice in a single line:
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
