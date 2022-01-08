package main

import "fmt"

func main() {
	// fmt.Println(avg([]float64{98, 93, 55, 33}))

	// fmt.Println(add(98, 93, 55, 33))
	// var r int
	// f(r)
	// fmt.Println(r)

	// x, y := f2()
	// fmt.Println(x, y)

	// add := func(x, y int) int {
	// 	return x + y
	// }

	// fmt.Println(add(1, 1))

	// defer schedule a function call at the end of a function
	// defer second()
	// first()
	// first()

	// panic and recover
	// pani is like throw in other programming languages
	// to handle it we have recover, but it used in this method

	// defer func() {
	// 	str := recover()
	// 	fmt.Println(str)
	// }()

	// panic("Panic")

	// pointers

	x := 5
	zero(&x)
	fmt.Println(x)

	// or

	yPtr := new(int)
	zero(yPtr)
	fmt.Println(yPtr)

}

func avg(slice []float64) float64 {
	//panic("Not implemented")

	total := 0.0
	for _, v := range slice {
		total += v
	}

	return total / float64(len(slice))
}

// variadic functions
func add(args ...int) int {
	total := 0

	for _, v := range args {
		total += v
	}

	return total
}

func f(r int) {
	r = 1
	return
}

func f2() (int, int) {
	return 4, 4
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func zero(xPtr *int) {
	*xPtr = 0
}
