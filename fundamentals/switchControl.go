package main

import "fmt"

func main(){
	fmt.Println("Input result:")
	var result int
	fmt.Scanln(&result)

	switch result {
	case 5:
		fmt.Println("Excelente")
	case 4:
		fmt.Println("Muy bien")
	case 3:
		fmt.Println("Bien")
	case 2:
		fmt.Println("Regular")
	case 1:
		fmt.Println("Reprobado")
	default:
		fmt.Println("Nota fuera del rango")
	}
}