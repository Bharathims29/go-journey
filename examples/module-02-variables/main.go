package main

import "fmt"

func main() {
	fmt.Println("Module 2: Variables, Types, and Operators")
	fmt.Println()

	var name string = "Bharathi"
	var age = 25
	city := "Madurai"
	const course = "Go Basic to Advanced"

	fmt.Println("Basic values")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	fmt.Println("Course:", course)
	fmt.Println()

	var emptyName string
	var emptyAge int
	var emptyPrice float64
	var isActive bool

	fmt.Println("Zero values")
	fmt.Println("String zero value:", emptyName)
	fmt.Println("Int zero value:", emptyAge)
	fmt.Println("Float zero value:", emptyPrice)
	fmt.Println("Bool zero value:", isActive)
	fmt.Println()

	marks := 95
	average := float64(marks)

	fmt.Println("Type conversion")
	fmt.Println("Marks as int:", marks)
	fmt.Println("Marks as float64:", average)
	fmt.Println()

	a := 10
	b := 3

	fmt.Println("Arithmetic operators")
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)
	fmt.Println()

	votingAge := 18
	hasID := true

	fmt.Println("Comparison and logical operators")
	fmt.Println("Is age greater than or equal to voting age?", age >= votingAge)
	fmt.Println("Can vote?", age >= votingAge && hasID)
	fmt.Println("Is minor?", age < votingAge)
	fmt.Println("Does not have ID?", !hasID)
}
