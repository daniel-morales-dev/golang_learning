package handlerErrors

import "fmt"

type InsufficientFunds struct {
	Amount  float64
	Balance float64
}

func (e *InsufficientFunds) Error() string {
	return fmt.Sprintf("No puedes transferir %.2f, tus fondos son insuficientes, tu balance: %.2f", e.Amount, e.Balance)
}
