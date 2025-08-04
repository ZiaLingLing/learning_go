package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type rectangle struct {
	length float64
	height float64
}

type triangle struct {
	base   float64
	height float64
}

func (e circle) areaCalc() float64 {
	var area float64 = math.Pi * float64(e.radius)
	return area
}

func (e rectangle) areaCalc() float64 {
	var area float64 = e.length * e.height
	return area
}

func (e triangle) areaCalc() float64 {
	var area float64 = e.base * e.height
	return area
}

type areaCalculator interface {
	areaCalc() float64
}

func calculateArea(e areaCalculator) float64 {
	return e.areaCalc()
}

func main() {
	var myCircle = circle{radius: 100}
	var myRectangle = rectangle{length: 10, height: 15}
	var myTriangle = triangle{base: 5, height: 10}

	fmt.Printf("Circle area: %.2f cm^2.\n", calculateArea(myCircle))
	fmt.Printf("Rectangle area: %.2f cm^2.\n", calculateArea(myRectangle))
	fmt.Printf("Triangle area: %.2f cm^2.\n", calculateArea(myTriangle))
}
