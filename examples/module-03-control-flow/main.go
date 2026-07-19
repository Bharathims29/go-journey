package main

import "fmt"

func main() {
	fmt.Println("Module 3: Control Flow")
	fmt.Println()

	age := 20

	fmt.Println("if and else")
	if age >= 18 {
		fmt.Println("Age", age, "means adult")
	} else {
		fmt.Println("Age", age, "means minor")
	}
	fmt.Println()

	marks := 82

	fmt.Println("else if")
	if marks >= 90 {
		fmt.Println("Grade: A")
	} else if marks >= 75 {
		fmt.Println("Grade: B")
	} else if marks >= 50 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Needs improvement")
	}
	fmt.Println()

	day := "Saturday"

	fmt.Println("switch")
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Regular weekday")
	}
	fmt.Println()

	level := 2

	fmt.Println("fallthrough")
	switch level {
	case 1:
		fmt.Println("Beginner")
	case 2:
		fmt.Println("Intermediate")
		fallthrough
	case 3:
		fmt.Println("Advanced practice unlocked")
	default:
		fmt.Println("Unknown level")
	}
	fmt.Println()

	fmt.Println("for loop")
	for i := 1; i <= 5; i++ {
		fmt.Println("Count:", i)
	}
	fmt.Println()

	fmt.Println("continue")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}

		fmt.Println("Printed number:", i)
	}
	fmt.Println()

	fmt.Println("break")
	for i := 1; i <= 10; i++ {
		if i == 6 {
			break
		}

		fmt.Println("Before break:", i)
	}
	fmt.Println()

	number := 19
	isPrime := true

	if number <= 1 {
		isPrime = false
	} else {
		for i := 2; i < number; i++ {
			if number%i == 0 {
				isPrime = false
				break
			}
		}
	}

	fmt.Println("Prime check")
	if isPrime {
		fmt.Println(number, "is prime")
	} else {
		fmt.Println(number, "is not prime")
	}
}
