package main

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func validateAge(age int) error {
	if age < 0 {
		return errors.New("age cannot be negative")
	}

	if age < 18 {
		return fmt.Errorf("age %d is below minimum age 18", age)
	}

	return nil
}

func findUser(name string) error {
	if name == "" {
		return ErrUserNotFound
	}

	return nil
}

func loadUser(name string) error {
	err := findUser(name)
	if err != nil {
		return fmt.Errorf("load user failed: %w", err)
	}

	return nil
}

func recoverExample() {
	defer func() {
		if value := recover(); value != nil {
			fmt.Println("Recovered from panic:", value)
		}
	}()

	panic("practice panic")
}

func main() {
	fmt.Println("Module 9: Error Handling")
	fmt.Println()

	fmt.Println("Successful division")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
	fmt.Println()

	fmt.Println("Division error")
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
	fmt.Println()

	fmt.Println("Validation error")
	err = validateAge(15)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println()

	fmt.Println("Wrapped error")
	err = loadUser("")
	if err != nil {
		fmt.Println("Error:", err)
	}

	if errors.Is(err, ErrUserNotFound) {
		fmt.Println("Checked with errors.Is: user was not found")
	}
	fmt.Println()

	fmt.Println("Panic and recover")
	recoverExample()
}
