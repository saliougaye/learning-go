package main

import "fmt"

func main() {

	// if

	if 10%2 == 0 {
		fmt.Println("Even")
	} else if 10/5 == 2 {
		fmt.Println("Divisible for 5")
	} else {
		fmt.Println("Odd")
	}

	// for as while in other programming language
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i += 1
	}

	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// switch case
	j := 5
	switch j {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	}

}
