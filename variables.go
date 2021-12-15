package main

import "fmt"

// global scope var declaration
var global_scope string = "Hi"

func main() {

	// declaration
	var s string

	// assignemnt
	s = "Hi Go"

	fmt.Println(s)

	// initialization
	var x string = "Hello World"

	fmt.Println(x)

	// shorthand declaration

	y := "Hello, World"

	fmt.Println(y)

	// scope
	fmt.Println(global_scope)

	// constants
	const z string = "Z Variable"

	// error
	//z = "ddssaddsa"

	// defining multiple variables
	var (
		a = 4
		b = 8
		c = 6
	)

	fmt.Println("a -> ", a, "b -> ", b, "c -> ", c)

	double()
}

func double() {
	fmt.Print("Enter number: ")
	var input float64

	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}
