package main

import "fmt"

// Son una lista, a diferencia de los arrays que son una caja cerrada, este es una caja abierta a la cual le podemos anexar mas items
func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice)

	mySlice2 := make([]string, 5)
	fmt.Println(mySlice2)

	var emptySlice []int
	fmt.Println(emptySlice)

	emptySlice = append(emptySlice, 20, 30, 40, 50)
	fmt.Println(emptySlice)

}
