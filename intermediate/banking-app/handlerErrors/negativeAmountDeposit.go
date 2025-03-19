package handlerErrors

import "fmt"

type NegativeAmountDeposit struct {
	Amount float64
}

func (e *NegativeAmountDeposit) Error() string {
	// Es el hermano fmt.Printf pero retorna un string formateado
	return fmt.Sprintf("No puedes depositar un saldo en negativo: %.2f", e.Amount)
}
