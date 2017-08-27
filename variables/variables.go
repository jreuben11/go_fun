package main

import "fmt"

func main() {
	// var declares 1 or more variables.
	var a string = "initial"
	fmt.Println(a)

	// can declare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)

	//  type inference of initialized variables.
	var d = true
	fmt.Println(d)

	// no initialization -> zero-valued
	var e int
	fmt.Println(e)

	// := syntax is shorthand for declaring and initializing a variable
	f := "short"
	fmt.Println(f)
}
