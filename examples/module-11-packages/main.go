package main

import (
	"fmt"

	"module11packages/calc"
	"module11packages/validate"
)

func main() {
	fmt.Println("Module 11: Packages, Modules, and Project Structure")
	fmt.Println()

	fmt.Println("Calculator package")
	fmt.Println("Add:", calc.Add(8, 4))

	result, ok := calc.Divide(8, 2)
	if ok {
		fmt.Println("Divide:", result)
	} else {
		fmt.Println("Cannot divide by zero")
	}
	fmt.Println()

	fmt.Println("Validation package")
	fmt.Println("Name is valid:", validate.Name("Bharath"))
	fmt.Println("Age is valid:", validate.Age(25))
}
