package main

import (
	"account/account"
	"account/handlerErrors"
	"errors"
	"fmt"
)

func main() {
	account1 := account.NewAccount("Daniel", 0)
	account2 := account.NewAccount("Nicole", 10)

	// En GO no existe el contexto de excepciones
	err1 := account1.Deposit(150)
	handleDepositErrors(err1)

	err2 := account2.Deposit(10)
	handleDepositErrors(err2)

	// Go desferencia la posición en memoria para ser mas amigable.
	fmt.Println("Account 1 ", account1) // Imprime &{Daniel 50}
	fmt.Println("Account 1 * ", *account1)
	fmt.Println("Account 1 & ", &account1) // Aquí seria como **account1, apuntaría a la posición en memoria

	printFinalBalance(account1)
	printFinalBalance(account2)

	errTransfer := account1.Transfer(30, account2)

	handleDepositErrors(errTransfer)

	printFinalBalance(account1)
	printFinalBalance(account2)

	errTransfer = account1.Transfer(-200, account2)
	handleDepositErrors(errTransfer)

}

func printFinalBalance(account *account.Account) {
	fmt.Printf("Dirección recibida en printFinalBalance: %p\n", account)
	fmt.Printf("El balance de la cuenta %s es %.2f\n", account.Owner, account.Balance)
}

func handleDepositErrors(err error) {
	if err != nil {
		// Desenrollando el error con errors.As
		var insufficientFundsErr *handlerErrors.InsufficientFunds
		var negativeDepositErr *handlerErrors.NegativeAmountDeposit

		if errors.As(err, &insufficientFundsErr) {
			fmt.Printf("ERROR: Fondos insuficientes - %s\n", insufficientFundsErr.Error())
		} else if errors.As(err, &negativeDepositErr) {
			fmt.Printf("ERROR: Depósito negativo - %s\n", negativeDepositErr.Error())
		} else {
			fmt.Printf("ERROR desconocido: %s\n", err.Error())
		}
		return
	}
}
