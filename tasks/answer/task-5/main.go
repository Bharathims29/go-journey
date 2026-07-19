package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Invalid deposit amount.")
		return
	}

	b.Balance += amount
	fmt.Printf("₹%.2f deposited successfully.\n", amount)
}

func (b *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Invalid withdrawal amount.")
		return
	}

	if amount > b.Balance {
		fmt.Println("Insufficient balance.")
		return
	}

	b.Balance -= amount
	fmt.Printf("₹%.2f withdrawn successfully.\n", amount)
}

func (b BankAccount) DisplayBalance() {
	fmt.Println("\n----- Account Details -----")
	fmt.Println("Account Holder :", b.Owner)
	fmt.Printf("Current Balance: ₹%.2f\n", b.Balance)
}

func main() {

	var account BankAccount

	fmt.Print("Enter Account Holder Name: ")
	fmt.Scan(&account.Owner)

	fmt.Print("Enter Initial Balance: ")
	fmt.Scan(&account.Balance)

	account.DisplayBalance()

	var depositAmount float64
	fmt.Print("\nEnter Deposit Amount: ")
	fmt.Scan(&depositAmount)
	account.Deposit(depositAmount)

	account.DisplayBalance()

	var withdrawAmount float64
	fmt.Print("\nEnter Withdrawal Amount: ")
	fmt.Scan(&withdrawAmount)
	account.Withdraw(withdrawAmount)

	account.DisplayBalance()
}
