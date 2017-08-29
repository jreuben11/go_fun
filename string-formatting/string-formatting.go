package main

import "fmt"
import "os"

type point struct {
	x, y int
}

func main() {
	//  Printf prints formatted string to os.Stdout.

	// prints an instance of  struct.
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	// include the struct’s field names
	fmt.Printf("%+v\n", p)

	// Go syntax representation of the value (source code snippet that would produce it)
	fmt.Printf("%#v\n", p)

	// type of a value
	fmt.Printf("%T\n", p)

	// booleans
	fmt.Printf("%t\n", true)

	// integer standard, base-10
	fmt.Printf("%d\n", 123)

	// integer binary representation.
	fmt.Printf("%b\n", 14)

	// character corresponding to the given integer.
	fmt.Printf("%c\n", 33)

	// hex encoding.
	fmt.Printf("%x\n", 456)

	// float basic decimals
	fmt.Printf("%f\n", 78.9)

	// float scientific notation.
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)

	// basic string printing
	fmt.Printf("%s\n", "\"string\"")

	// double-quote strings as in Go source
	fmt.Printf("%q\n", "\"string\"")

	// render string in base-16, with two output characters per byte of input.
	fmt.Printf("%x\n", "hex this")

	// print a representation of a pointer
	fmt.Printf("%p\n", &p)

	// integer width - use a number after the %. default is right-justified and padded with spaces.
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	// float width and precision
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	// To left-justify, use -
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// string width  --> align in table-like output. For basic right-justified width.
	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	// string left-justify
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// Sprintf formats and returns a string without printing
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// Fprintf: format+print to io.Writers other than os.Stdout
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
