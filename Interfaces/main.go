package main

import (
	"fmt"
	"math"
	"reflect"
)

type shape interface {
	calculateArea() float64
	calculatePerimeter() float64
}

type circle struct {
	radius float64
}

func (c circle) calculateArea() float64 {
	diameter := math.Pow(c.radius, 2)
	return diameter * math.Pi
}

func (c circle) calculatePerimeter() float64 {
	return 2 * math.Pi * c.radius
}

type rectangle struct {
	length float64
	width  float64
}

func (r rectangle) calculateArea() float64 {
	return r.width * r.length
}

func (r rectangle) calculatePerimeter() float64 {
	return 2 * (r.width + r.length)
}

func main() {

	circle1 := circle{
		radius: 2.5,
	}

	rectangle1 := rectangle{
		length: 3,
		width:  2,
	}

	shapes := []shape{circle1, rectangle1}
	for _, shape := range shapes {
		perimeter := shape.calculatePerimeter()
		area := shape.calculateArea()

		fmt.Println("Perimeter of", reflect.TypeOf(shape), "is", perimeter)
		fmt.Println("Area of", reflect.TypeOf(shape),"is", area)
	}

}
