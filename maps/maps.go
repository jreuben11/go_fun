package main

import "fmt"

func main() {
	// To create an empty map, use  make
	m := make(map[string]int)

	// Set key/value pairs
	m["k1"] = 7
	m["k2"] = 13

	// Printing a map
	fmt.Println("map:", m)

	// Get
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// builtin len
	fmt.Println("len:", len(m))

	// builtin delete removes key/value pairs from a map.
	delete(m, "k2")
	fmt.Println("map:", m)

	// optional second return value when getting -  disambiguate between missing keys and keys with zero values
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// declare and initialize a new map in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
