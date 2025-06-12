package basics

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition: ", result)

	result = a - b
	fmt.Println("Subtraction: ", result)

	result = a * b
	fmt.Println("Multiplication: ", result)

	result = a / b
	fmt.Println("Division: ", result)

	result = a % b
	fmt.Println("Remainder: ", result)

	const p float64 = 22 / 7.0
	fmt.Println(p)

	// Overflow with signed Integer
	var maxInt int64 = 9223372036854775807 // Max for signed int64 can hold goes from -9223372036854775807 to 9223372036854775807
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

	// Overflow with unsigned Integer
	var uMaxInt uint64 = 18446744073709551615 // Max for unsigned in64 goes from 0 to 18446744073709551615
	fmt.Println(uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64 // Too small to the point where it underflow to 0
	fmt.Println(smallFloat)
}
