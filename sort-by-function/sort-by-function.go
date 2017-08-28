package main

import "sort"
import "fmt"

// need a type alias
type ByLength []string

// implement sort.Interface - Len, Less, and Swap - on the type -> can use the sort package’s generic Sort function
func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// can now implement custom sort by casting a slice to the type, and then use sort.Sort on that typed slice.
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}
