package main

import (
	"fmt"
)

// Define the interface for account operations
type AccountOperations interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	Balance() float64
}

// Define the BankAccount struct that implements the AccountOperations interface
type BankAccount struct {
	accountNumber string
	accountHolder string
	balance       float64
}

// Implement methods for BankAccount to satisfy the AccountOperations interface
func (a *BankAccount) Deposit(amount float64) {
	a.balance += amount
	fmt.Printf("Deposited %.2f. Current balance: %.2f\n", amount, a.balance)
}

func (a *BankAccount) Withdraw(amount float64) error {
	if a.balance >= amount {
		a.balance -= amount
		fmt.Printf("Withdrawn %.2f. Current balance: %.2f\n", amount, a.balance)
		return nil
	}
	return fmt.Errorf("insufficient funds. Current balance: %.2f", a.balance)
}

func (a *BankAccount) Balance() float64 {
	return a.balance
}

func BankAction() {
	// Create a new BankAccount instance
	account := &BankAccount{
		accountNumber: "123456789",
		accountHolder: "John Doe",
		balance:       1000.00,
	}

	// Perform account operations using the interface
	account.Deposit(500.00)
	account.Withdraw(200.00)
	account.Withdraw(1500.00) // This should result in an error due to insufficient funds

	// Print the final balance
	fmt.Printf("Final balance for account %s: %.2f\n", account.accountNumber, account.Balance())
}
