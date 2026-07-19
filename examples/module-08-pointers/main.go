package main

import "fmt"

type Student struct {
	Name  string
	Marks float64
}

type BankAccount struct {
	Owner   string
	Balance float64
}

func changeWithoutPointer(number int) {
	number = 100
}

func changeWithPointer(number *int) {
	*number = 100
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func updateMarks(student *Student, marks float64) {
	student.Marks = marks
}

func (account *BankAccount) Deposit(amount float64) {
	account.Balance += amount
}

func main() {
	fmt.Println("Module 8: Pointers")
	fmt.Println()

	age := 25
	agePointer := &age

	fmt.Println("Address and dereference")
	fmt.Println("Age value:", age)
	fmt.Println("Age address:", agePointer)
	fmt.Println("Value through pointer:", *agePointer)

	*agePointer = 30
	fmt.Println("Age after pointer update:", age)
	fmt.Println()

	number := 10

	fmt.Println("Pass by value")
	changeWithoutPointer(number)
	fmt.Println("After changeWithoutPointer:", number)

	fmt.Println("Pointer parameter")
	changeWithPointer(&number)
	fmt.Println("After changeWithPointer:", number)
	fmt.Println()

	x := 10
	y := 20

	fmt.Println("Swap with pointers")
	fmt.Println("Before swap:", x, y)
	swap(&x, &y)
	fmt.Println("After swap:", x, y)
	fmt.Println()

	student := Student{
		Name:  "Bharath",
		Marks: 80,
	}

	fmt.Println("Struct pointer")
	fmt.Println("Before marks update:", student)
	updateMarks(&student, 92.5)
	fmt.Println("After marks update:", student)
	fmt.Println()

	account := BankAccount{
		Owner:   "Bharath",
		Balance: 1000,
	}

	fmt.Println("Pointer receiver")
	fmt.Println("Before deposit:", account.Balance)
	account.Deposit(500)
	fmt.Println("After deposit:", account.Balance)
}
