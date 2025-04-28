package trainingErrors

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func withdraw(balance, amount float64) (float64, error) {
	if amount > balance {
		return balance, fmt.Errorf("insufficient funds: trying to withdraw %.2f, but balance is %.2f", amount, balance)
	}
	return balance - amount, nil
}

// Errores personalizados

type InsufficientFundsError struct {
	Balance float64
	Attempt float64
}

func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("Saldo insuficiente: balance=%.2f, intento=%.2f", e.Balance, e.Attempt)
}

func withdrawV2(balance, amount float64) (float64, error) {
	if amount > balance {
		return balance, &InsufficientFundsError{Balance: balance, Attempt: amount}
	}
	return balance - amount, nil
}
