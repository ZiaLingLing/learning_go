package main

import (
	"fmt"
	"learning_go/internal/shapeareacalc"
)

func main() {
	var myCircle = shapeareacalc.Circle{Radius: 100}
	var myRectangle = shapeareacalc.Rectangle{Length: 10, Height: 15}
	var myTriangle = shapeareacalc.Triangle{Base: 5, Height: 10}

	fmt.Printf("Circle area: %.2f cm^2.\n", shapeareacalc.CalculateArea(myCircle))
	fmt.Printf("Rectangle area: %.2f cm^2.\n", shapeareacalc.CalculateArea(myRectangle))
	fmt.Printf("Triangle area: %.2f cm^2.\n", shapeareacalc.CalculateArea(myTriangle))
}
