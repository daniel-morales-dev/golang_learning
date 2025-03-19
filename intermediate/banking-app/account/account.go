package account

import (
	"account/handlerErrors"
	"fmt"
)

// Account represents a simple bank account.
type Account struct {
	Owner   string
	Balance float64
}

// Esto es el equivalente a un constructor
// Retorna la posición en memoria

// NewAccount creates a new Account with the given owner and balance.
func NewAccount(owner string, balance float64) *Account {
	return &Account{Owner: owner, Balance: balance}
}

func (account *Account) Deposit(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("error en depósito: %w", &handlerErrors.NegativeAmountDeposit{Amount: amount})
	}
	account.Balance += amount
	fmt.Printf("Deposited %.2f to %s's account. New balance: %.2f\n", amount, account.Owner, account.Balance)
	return nil
}

func (account *Account) Transfer(amount float64, recipient *Account) error {
	fmt.Printf("Try to transfer %.2f to %s's account.\n", amount, recipient.Owner)

	if amount >= account.Balance {
		return fmt.Errorf("error en transferencia: %w", &handlerErrors.InsufficientFunds{Amount: amount, Balance: account.Balance})
	}

	err := recipient.Deposit(amount)
	if err != nil {
		return fmt.Errorf("falló depósito en cuenta destino: %w", err)
	}
	account.Balance -= amount

	fmt.Printf("Transferred %.2f to %s's account. New balance: %.2f\n", amount, recipient.Owner, recipient.Balance)

	return nil
}
