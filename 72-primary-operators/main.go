package main

import (
	"fmt"
	"math"
)

func main() {
	// multiplication
	var radius = 12.0

	area := math.Pi * radius * radius
	fmt.Println("area", area)

	// integer division
	half := 1 / 2
	fmt.Println("Half", half)

	halfFloat := 1.0 / 2.0
	fmt.Println("half float", halfFloat)

	// squaring, and raising to power
	badThreeSquared := 3 ^ 2 // this is the bitwise XOR operator
	fmt.Println("Three squared is not", badThreeSquared)
	goodThreeSquared := math.Pow(3.0, 2.0)
	fmt.Println("Three squared is", goodThreeSquared)

	// modulus operator
	remainder := 50 % 3
	fmt.Println("remaider is", remainder)

	// unary operators
	x := 3
	// can't do this
	//y := x++
	// can do this
	var y = x
	y++
	fmt.Println("x is", x, "and y is", y)
}
