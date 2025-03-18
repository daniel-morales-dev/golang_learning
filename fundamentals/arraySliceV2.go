package main

import "fmt"

func main() {
	var listNumber [5]int
	fmt.Println("Ingresa 5 números separados por espacios")
	for i := 0; i < 5; i++ {
		fmt.Scan(&listNumber[i])
	}

	slice := listNumber[:]

	fmt.Println("Lista de números: ", slice)

	sum := 0
	for _, value := range slice {
		sum += value
	}

	fmt.Printf("La suma del slice es %v\n", sum)

}
