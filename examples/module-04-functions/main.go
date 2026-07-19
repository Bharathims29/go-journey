package main

import "fmt"

func sayWelcome() {
	fmt.Println("Welcome to Module 4: Functions")
}

func greetUser(name string) {
	fmt.Println("Hello,", name)
}

func add(a int, b int) int {
	return a + b
}

func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b

	return quotient, remainder
}

func rectangle(width int, height int) (area int, perimeter int) {
	area = width * height
	perimeter = 2 * (width + height)
	return
}

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func showScope() {
	message := "This variable exists only inside showScope"
	fmt.Println(message)
}

func main() {
	sayWelcome()
	fmt.Println()

	greetUser("Bharath")
	fmt.Println()

	total := add(10, 20)
	fmt.Println("Add result:", total)
	fmt.Println()

	quotient, remainder := divide(10, 3)
	fmt.Println("Divide result")
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)
	fmt.Println()

	area, perimeter := rectangle(5, 3)
	fmt.Println("Rectangle result")
	fmt.Println("Area:", area)
	fmt.Println("Perimeter:", perimeter)
	fmt.Println()

	fmt.Println("Variadic function")
	fmt.Println("Sum:", sum(10, 20, 30, 40, 50))
	fmt.Println()

	fmt.Println("Scope")
	showScope()
	fmt.Println()

	double := func(number int) int {
		return number * 2
	}

	fmt.Println("Anonymous function")
	fmt.Println("Double of 5:", double(5))
}
