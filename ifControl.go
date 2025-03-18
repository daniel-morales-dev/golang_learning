package main

import "fmt"

func main(){
	fmt.Println("Input a number:")
	var number int8
	fmt.Scanln(&number)

	fmt.Println(examineNumber(number))
}

func examineNumber(number int8) string{
	if number > 0 {
		return "El numero es positivo"
	}else if number <0{
		return "El numero es negativo"
	}

	return "El numero es 0"
}