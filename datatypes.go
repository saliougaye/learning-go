package main

import "fmt"

func main() {

	// numbers

	// integers

	fmt.Println("5 + 2 =", 5+2)

	// float
	fmt.Println("8.8 + 2.2 =", 8.8+2.1)

	// strings

	fmt.Println("Hi Go!")

	// len return string length
	fmt.Println(len("Hi Go!"))

	// return character of a string
	fmt.Println("Hi Go!"[3])

	// concatenate string
	fmt.Println("Hi " + "Go!")

	// booleans
	fmt.Println("True: ", true, ", False: ", false)

	// booleans conditions
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}
