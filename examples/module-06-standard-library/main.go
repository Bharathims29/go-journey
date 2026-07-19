package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("Module 6: Strings and Basic Standard Library")
	fmt.Println()

	name := "Bharath"
	message := "Hello, " + name

	fmt.Println("Basic strings")
	fmt.Println("Name:", name)
	fmt.Println("Message:", message)
	fmt.Println("Name byte length:", len(name))
	fmt.Println()

	text := "  Go is simple, Go is powerful  "

	fmt.Println("strings package")
	fmt.Println("Original:", text)
	fmt.Println("Trimmed:", strings.TrimSpace(text))
	fmt.Println("Uppercase:", strings.ToUpper(text))
	fmt.Println("Lowercase:", strings.ToLower(text))
	fmt.Println("Contains Go?", strings.Contains(text, "Go"))
	fmt.Println("Replace:", strings.ReplaceAll(text, "powerful", "productive"))
	fmt.Println("Words:", strings.Fields(text))
	fmt.Println("Split:", strings.Split("go,java,python", ","))
	fmt.Println("Join:", strings.Join([]string{"go", "java", "python"}, " | "))
	fmt.Println()

	fmt.Println("strconv package")
	age, err := strconv.Atoi("25")
	if err != nil {
		fmt.Println("Age conversion failed:", err)
	} else {
		fmt.Println("Age plus 5:", age+5)
	}

	price, err := strconv.ParseFloat("99.50", 64)
	if err != nil {
		fmt.Println("Price conversion failed:", err)
	} else {
		fmt.Println("Price with tax:", price+10.50)
	}

	ageText := strconv.Itoa(25)
	fmt.Println("Age as string:", ageText)
	fmt.Println()

	fmt.Println("conversion error example")
	_, err = strconv.Atoi("not-a-number")
	if err != nil {
		fmt.Println("Could not convert value:", err)
	}
	fmt.Println()

	fmt.Println("math package")
	fmt.Println("Square root of 64:", math.Sqrt(64))
	fmt.Println("2 power 3:", math.Pow(2, 3))
	fmt.Println("Round 10.6:", math.Round(10.6))
	fmt.Println("Ceil 10.2:", math.Ceil(10.2))
	fmt.Println("Floor 10.8:", math.Floor(10.8))
	fmt.Println()

	now := time.Now()

	fmt.Println("time package")
	fmt.Println("Current time:", now)
	fmt.Println("Date:", now.Format("2006-01-02"))
	fmt.Println("Time:", now.Format("15:04:05"))
	fmt.Println("Time with Zone:", now.UTC())
	fmt.Println()

	username := "  bharath_go  "
	cleanUsername := strings.TrimSpace(username)

	fmt.Println("Username validation")
	if len(cleanUsername) >= 3 && !strings.Contains(cleanUsername, " ") {
		fmt.Println("Valid username:", cleanUsername)
	} else {
		fmt.Println("Invalid username")
	}
}
