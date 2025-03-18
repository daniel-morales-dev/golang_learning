package handlerErrors

import "fmt"

type NegativeAmountDeposite struct {
	Amount float64
}

func (e *NegativeAmountDeposite) Error() string {
	// Es el hermano fmt.Printf pero retorna un string formateado
	return fmt.Sprintf("No puedes depositar un saldo en negativo: %.2f", e.Amount)
}
