package basics

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	//pascal case  PascalCase
	//Structs, Interface, emums

	// Snake case  snake_case
	// Name files in snake case, variables and functions in camel case

	//UPPERCASE  UPPERCASE
	// Constants in uppercase

	//mixed case  mixedCase
	// Variables and functions in camel case isValid, getName, setName

	const PI = 3.14
	var employeeId = 12345
	fmt.Println("PI:", PI)
	fmt.Println("Employee ID:", employeeId)

}