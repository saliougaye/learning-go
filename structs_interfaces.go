package main

import (
	"fmt"
	"math"
)

// use of struct
// declaration of a struct
// type name struct { fields }
type Circle struct {
	x, y, r float64
}

// its a method, not a function and we call it from a Circle type variable
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y2)

	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1

	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y2)

	return l * w
}

// func circleArea(x, y, r float64) float64 {
// 	return math.Pi * r * r
// }

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("hi, my name is ", p.Name)
}

type Android struct {
	Person Person
	Model  string
}

// Rectangle and Circle have area method, so they implemented this interface
type Shape interface {
	area() float64
}

// variadic function that calculate area for all shapes

func totalArea(shapes ...Shape) float64 {
	var area float64

	for _, s := range shapes {
		area += s.area()
	}

	return area
}

func main() {
	// program to calculate the area of a rectangle and a circle

	// var rx1, ry1 float64 = 0, 0
	// var rx2, ry2 float64 = 10, 10
	// var cx, cy, cr float64 = 0, 0, 5
	// fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	// fmt.Println(circleArea(cx, cy, cr))

	// declaration of a struct
	//var c Circle

	// or
	//c := new(Circle)

	// declaration and initalization
	c := Circle{
		x: 0,
		y: 0,
		r: 5,
	}

	// or
	// c := Circle{0, 0, 5}
	//or with pointer
	// c := &Circle{0, 0, 5}

	// edit fields
	c.x = 10
	c.y = 5
	fmt.Println(c.x, c.y, c.r)

	fmt.Println("[Function] Area Circle: ", circleArea(&c))
	fmt.Println("[Method] Area Circle: ", c.area())

	r := Rectangle{
		x1: 0,
		y1: 0,
		x2: 10,
		y2: 10,
	}

	fmt.Println("[Method] Area Rectangle: ", r.area())

	// nested struct
	a := new(Android)
	a.Person.Talk()

	fmt.Println(totalArea(&c, &r))

}
