package main

import (
	"fmt"
	"strings" // contains strings method
)

func main() {

	// contains
	fmt.Println(strings.Contains("test", "es"))

	// count letter
	fmt.Println(strings.Count("test", "t"))

	// hasprefix: if a smaller string starts in a bigger string
	fmt.Println(strings.HasPrefix("test", "te"))

	// hasprefix: if a bigger string ends with a smaller string
	fmt.Println(strings.HasPrefix("test", "st"))

	// find position of a char in a string
	fmt.Println(strings.Index("test", "e"))

	// join a list of strings and separate with a specific char
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))

	// repeat a string
	fmt.Println(strings.Repeat("test", 3))

	// replace
	fmt.Println(strings.Replace("aaaa", "a", "b", 3))

	// split
	fmt.Println(strings.Split("a-b-c-d-e", "-"))

	// lowercase
	fmt.Println(strings.ToLower("TEST"))

	//uppercase
	fmt.Println(strings.ToUpper("test"))
}
