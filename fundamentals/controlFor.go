package main

import "fmt"

func main() {
	fmt.Println("Input a number:")
	var number int
	fmt.Scan(&number)
	fmt.Printf("Tabla de multiplicar del %v\n", number)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%vx%v=%v\n", i, number, i*number)
	}
}
