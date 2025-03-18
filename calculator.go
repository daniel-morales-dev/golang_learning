package main

import (
	"fmt"
)

func main() {
	var firstNumber float64
	var secondNumber float64
	var operator string
	var exit int = 1

	for exit != 0 {
		fmt.Println("Calculadora")
		fmt.Println("Input first number:")
		fmt.Scanln(&firstNumber)
		fmt.Println("Input second number:")
		fmt.Scanln(&secondNumber)
		fmt.Println("Input operator:")
		fmt.Scanln(&operator)

		switch operator {
		case "+":
			fmt.Println("Operación de suma")
			fmt.Println(firstNumber + secondNumber)
		case "-":
			fmt.Println("Operación de resta")
			fmt.Println(firstNumber - secondNumber)
		case "*":
			fmt.Println("Operación de multiplicación")
			fmt.Println(firstNumber * secondNumber)
		case "/":
			fmt.Println("Operación de división")
			if secondNumber == 0 {
				fmt.Println("No se puede dividir por cero")
			}
			fmt.Println(firstNumber / secondNumber)
		default:
			fmt.Println("Operador no válido")
		}

		fmt.Println("Deseas salir? (0 para salir, 1 para continuar)")
		fmt.Scanln(&exit)
	}
	fmt.Println("Adios!")
}
