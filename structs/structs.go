package main

import "fmt"

// struct type
type person struct {
	name string
	age  int
}

func main() {

	// create a new struct
	fmt.Println(person{"Bob", 20})

	// can name  fields when initializing
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"})

	// An & prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})

	// Access struct fields with a dot.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// can also use dots with struct pointers - the pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age)
}
