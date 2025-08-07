package shapeareacalc

import "math"

func (e Circle) AreaCalc() float64 {
	return math.Pi * e.Radius * e.Radius
}

func (e Rectangle) AreaCalc() float64 {
	return e.Length * e.Height
}

func (e Triangle) AreaCalc() float64 {
	return 0.5 * e.Base * e.Height
}

type AreaCalculator interface {
	AreaCalc() float64
}

func CalculateArea(e AreaCalculator) float64 {
	return e.AreaCalc()
}
