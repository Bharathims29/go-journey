package main

import "fmt"

func checkNumber(num int) string {
	if num > 0 {
		return "Positive"
	} else if num < 0 {
		return "Negative"
	}
	return "Zero"
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
func isEven(num int) bool {
	return num%2 == 0
}

func analyzeNumber(num int) (string, bool, bool) {
	numberType := checkNumber(num)
	prime := isPrime(num)
	even := isEven(num)
	return numberType, prime, even
}

func main() {
	var num int

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	numberType, prime, even := analyzeNumber(num)

	fmt.Println("Number Type:", numberType)
	if prime {
		fmt.Println("Prime Number")
	} else {
		fmt.Println("Not a Prime Number")
	}
	if even {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}

}
