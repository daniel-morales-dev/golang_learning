package main

import (
	"account/account"
	"fmt"
)

func main() {
	account1 := account.NewAccount("Daniel", 0)
	account2 := account.NewAccount("Nicole", 10)

	// En GO no existe el contexto de excepciones
	err1 := account1.Deposit(100)
	if err1 != nil {
		fmt.Println(err1)
		//log.Fatal(err1) // Termina y muestra el error
		return // Salir porque no tiene sentido seguir con un depósito inválido
	}
	err2 := account2.Deposit(10)
	if err2 != nil {
		fmt.Println(err2)
		//log.Fatal(err2) // Termina y muestra el error
		return // Salir porque no tiene sentido seguir con un depósito inválido
	}

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

func printFinalBalance(account *account.Account) {
	fmt.Printf("Dirección recibida en printFinalBalance: %p\n", account)
	fmt.Printf("El balance de la cuenta %s es %.2f\n", account.Owner, account.Balance)
}
