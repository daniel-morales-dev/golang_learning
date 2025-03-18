package main

import "fmt"

func main(){
	var number1 int = 10
	var number2 int8 = 20
	var number3 int16 = 30
	var number4 int32 = 40
	var number5 int64 = 50
	var number6 uint = 60
	var number7 uint8 = 70
	var number8 uint16 = 80
	var number9 uint32 = 90
	var number10 uint64 = 100

	fmt.Printf("Valor y tipo numero 1: %v, %T\n", number1, number1)
	fmt.Printf("Valor y tipo numero 2: %v, %T\n", number2, number2)
	fmt.Printf("Valor y tipo numero 3: %v, %T\n", number3, number3)
	fmt.Printf("Valor y tipo numero 4: %v, %T\n", number4, number4)
	fmt.Printf("Valor y tipo numero 5: %v, %T\n", number5, number5)
	fmt.Printf("Valor y tipo numero 6: %v, %T\n", number6, number6)
	fmt.Printf("Valor y tipo numero 7: %v, %T\n", number7, number7)
	fmt.Printf("Valor y tipo numero 8: %v, %T\n", number8, number8)
	fmt.Printf("Valor y tipo numero 9: %v, %T\n", number9, number9)
	fmt.Printf("Valor y tipo numero 10: %v, %T\n", number10, number10)

	var number11 float64 = 10.1
	var number12 int = 10

	fmt.Printf("Suma de numero 11 y numero 12: %v", number11 + float64(number12))
}