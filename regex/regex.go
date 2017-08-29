package main

import "bytes"
import "fmt"
import "regexp"

func main() {
	// use a string pattern directly - test whether a pattern matches
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Compile an optimized Regexp struct.
	r, _ := regexp.Compile("p([a-z]+)ch")
	// match test
	fmt.Println(r.MatchString("peach"))

	// find
	fmt.Println(r.FindString("peach punch"))

	// find first match & return start and end indexes
	fmt.Println(r.FindStringIndex("peach punch"))

	//  return information for both p([a-z]+)ch and ([a-z]+).
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// return information about the indexes of matches and submatches.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// All variants - all matches in the input, not just first
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// limit the number of matches - provide non-negative integer as the second argument
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// use []byte arguments - drop String from the function name.
	fmt.Println(r.Match([]byte("peach")))

	// constants with regular expressions - use MustCompile
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// replace subsets of strings
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// Func variant - transform matched text with a given function.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
