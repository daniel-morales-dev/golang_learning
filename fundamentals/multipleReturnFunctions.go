package main

import "fmt"

func main(){
	var number1 int = 10
	var number2 int = 20

	sum, difference, product, quotient := calculator(number1, number2)

	fmt.Printf("Suma: %v\nResta: %v\nMultiplicación: %v\nDivision: %v",sum, difference, product, quotient)
}

func calculator(number1 int, number2 int)(int, int, int, float64){
	sum := number1 + number2
	difference := number1 - number2
	product := number1 * number2
	quotient := float64(number1) / float64(number2)

	return sum, difference, product, quotient
}