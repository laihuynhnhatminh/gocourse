package basics

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase
	// Example: CalculateArea, GetUserInfo
	// Structs, interfaces, enums

	// snake_case
	// Example: user_id, first_name, get_request

	// UPPERCASE
	// Example: HOUR, MINUTE

	// mixedCase
	// Example: javascriptApplication, htmlDocument, isValid

	const MAXRETRIES = 5

	var employeeID = 1001
	fmt.Print("EmployeeID: ", employeeID)
}
