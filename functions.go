package main

import "fmt"

func main() {
	fmt.Println(avg([]float64{98, 93, 55, 33}))

	fmt.Println(add(98, 93, 55, 33))
	var r int
	f(r)
	fmt.Println(r)

	x, y := f2()
	fmt.Println(x, y)
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
