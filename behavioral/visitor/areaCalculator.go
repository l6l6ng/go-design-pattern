package main

/*具体访问者*/

import "fmt"

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	fmt.Println("Calculating area for square")
}

func (a *areaCalculator) visitForCircle(s *circle) {
	fmt.Println("Calculating area for circle")
}
func (a *areaCalculator) visitForRectangle(s *rectangle) {
	fmt.Println("Calculating area for rectangle")
}