package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Marks float64
}

func (s Student) Display() {
	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)
	fmt.Println("Marks:", s.Marks)
}

func (s Student) IsPassed() bool {
	return s.Marks >= 40
}

func (s *Student) UpdateMarks(marks float64) {
	s.Marks = marks
}

type Address struct {
	City    string
	Country string
}

type Employee struct {
	Name    string
	Role    string
	Address Address
}

type BankAccount struct {
	Owner   string
	Balance float64
}

func (a BankAccount) DisplayBalance() {
	fmt.Println(a.Owner, "balance:", a.Balance)
}

func (a *BankAccount) Deposit(amount float64) {
	a.Balance += amount
}

func (a *BankAccount) Withdraw(amount float64) bool {
	if amount > a.Balance {
		return false
	}

	a.Balance -= amount
	return true
}

func main() {
	fmt.Println("Module 7: Structs and Methods")
	fmt.Println()

	student := Student{
		Name:  "Bharath",
		Age:   25,
		Marks: 88.5,
	}

	fmt.Println("Student struct")
	student.Display()
	fmt.Println("Passed?", student.IsPassed())
	fmt.Println()

	student.UpdateMarks(92.0)

	fmt.Println("After updating marks")
	student.Display()
	fmt.Println()

	employee := Employee{
		Name: "Anu",
		Role: "Developer",
		Address: Address{
			City:    "Bengaluru",
			Country: "India",
		},
	}

	fmt.Println("Nested struct")
	fmt.Println("Employee:", employee.Name)
	fmt.Println("Role:", employee.Role)
	fmt.Println("City:", employee.Address.City)
	fmt.Println("Country:", employee.Address.Country)
	fmt.Println()

	account := BankAccount{
		Owner:   "Bharath",
		Balance: 1000,
	}

	fmt.Println("Bank account")
	account.DisplayBalance()

	account.Deposit(500)
	account.DisplayBalance()

	if account.Withdraw(300) {
		fmt.Println("Withdraw successful")
	} else {
		fmt.Println("Withdraw failed")
	}

	account.DisplayBalance()
}
