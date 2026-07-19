package main

import (
	"errors"
	"fmt"
	"strings"
)

func Add(a, b int) int {
	return a + b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func IsValidUsername(username string) bool {
	if len(username) < 3 {
		return false
	}

	for _, ch := range username {
		if ch == ' ' {
			return false
		}
	}

	return true
}

func NormalizeName(name string) string {
	return strings.ToLower(strings.TrimSpace(name))
}

func main() {
	fmt.Println(Add(3, 4))
	fmt.Println(Divide(10, 2))
	fmt.Println(IsValidUsername("bharathi"))
	fmt.Println(NormalizeName("bharathi "))
}
