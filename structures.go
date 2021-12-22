package main

import "fmt"

func main() {
	// arrays
	var x [5]int
	x[4] = 100

	fmt.Println(x)

	//slices

	// declaration
	var y []float64

	// declare a slice with an underlying float64 arrys of length 5
	y = make([]float64, 5)

	fmt.Println(y)

	z := x[2:5]

	fmt.Println(z)

	// append to slice
	y = append(y, 1, 2)

	//copy
	slice1 := []int{1, 2, 3, 4, 6}
	slice2 := make([]int, 4)
	copy(slice2, slice1)

	fmt.Println(slice1, slice2)

	// maps

	// in this example, we cannot use the map in this mode because is not initialized
	// var m map[string]int

	// m["key"] = 10

	// fmt.Println(m)

	m := make(map[string]int)

	m["key"] = 10

	fmt.Println(m)

	// delete from a map
	// delete(m, "key")

	// the second parameter tells if the lookup in the map is successful
	// if the key not exist, the ok variable is false, else is true
	name, ok := m["key"]

	fmt.Println(name, ok)

}
