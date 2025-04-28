package main

import (
	"errors"
	"fmt"
)

type Account struct {
	Owner   string
	Balance float64
}

// Esto es el equivalente a un constructor
// Retorna la posición en memoria
func NewAccount(owner string, balance float64) *Account {
	return &Account{Owner: owner, Balance: balance}
}

func (account *Account) Transfer(amount float64, recipient *Account) error {
	fmt.Printf("Try to transfer %.2f to %s's account.\n", amount, recipient.Owner)

	if amount >= account.Balance {
		return errors.New("No tienes el monto suficiente para esta transferencia.")
	}

	recipient.Deposit(amount)
	account.Balance -= amount

	fmt.Printf("Transferred %.2f to %s's account. New balance: %.2f\n", amount, recipient.Owner, recipient.Balance)

	return nil
}

func (account *Account) Deposit(amount float64) {
	account.Balance += amount
	fmt.Printf("Deposited %.2f to %s's account. New balance: %.2f\n", amount, account.Owner, account.Balance)
}

func main() {
	account1 := NewAccount("Daniel", 0)
	account2 := NewAccount("Nicole", 10)

	account1.Deposit(50)
	account2.Deposit(10)

	// Go desferencia la posición en memoria para ser mas amigable.
	fmt.Println("Account 1 ", account1) // Imprime &{Daniel 50}
	fmt.Println("Account 1 * ", *account1)
	fmt.Println("Account 1 & ", &account1) // Aquí seria como **account1, apuntaría a la posición en memoria

	printFinalBalance(account1)
	printFinalBalance(account2)

	account1.Transfer(30, account2)

	printFinalBalance(account1)
	printFinalBalance(account2)

	err := account1.Transfer(100, account2)
	if err != nil {
		fmt.Println(err)
	}
}

// Si le pasamos una posicion en memoria, Go se encarga de desreferenciarla
func printFinalBalance(account *Account) {
	fmt.Printf("Dirección recibida en printFinalBalance: %p\n", account)
	fmt.Printf("El balance de la cuenta %s es %.2f\n", account.Owner, account.Balance)
}
