package basic

import "fmt"

var middleName = "Nhat"

func main() {
	var age int
	var firstName string = "Minh"

	count := 10
	lastName := "Lai"

	// Default Value
	// Numeric: 0
	// Boolean: false
	// String: ""
	// Pointers, slices, maps, functions and structs: nil

	// ~~~~ SCOPE

	fmt.Println(age)
	fmt.Println(firstName)
	fmt.Println(count)
	fmt.Println(lastName)
	fmt.Println(middleName)
}

func printname() {
	firstName := "Minh"
	fmt.Println(firstName)
}
