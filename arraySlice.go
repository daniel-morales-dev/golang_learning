package main

import "fmt"

func main() {
	var listNumber [5]int
	fmt.Println("Ingresa 5 números separados por espacios")
	var n1, n2, n3, n4, n5 int

	fmt.Scanln(&n1, &n2, &n3, &n4, &n5)

	listNumber[0] = n1
	listNumber[1] = n2
	listNumber[2] = n3
	listNumber[3] = n4
	listNumber[4] = n5

	slice := listNumber[:]

	fmt.Println("Lista de números: ", slice)

	sum := 0
	for _, value := range slice {
		sum += value
	}

	fmt.Printf("La suma del slice es %v\n", sum)

}
