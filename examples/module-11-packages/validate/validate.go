package validate

import "strings"

func Name(name string) bool {
	return len(strings.TrimSpace(name)) >= 3
}

func Age(age int) bool {
	return isPositive(age) && age >= 18
}

func isPositive(number int) bool {
	return number > 0
}
