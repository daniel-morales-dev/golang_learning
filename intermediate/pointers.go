package main

import "fmt"

/*
Un puntero es una variable que almacena la dirección de memoria de otra variable.
Es decir, en vez de guardar un valor directo, guarda dónde está el valor en memoria.
*/
func main() {
	var x int = 10
	var p *int = &x // p es un puntero que apunta a x

	fmt.Println(x)  // 10
	fmt.Println(p)  // dirección de memoria de x (ej: 0xc000014088)
	fmt.Println(*p) // 10 -> desreferenciando, accedemos al valor al que apunta

	number := 5
	increment(&number)
	fmt.Println(number) // 6
}

func increment(val *int) {
	*val = *val + 1
}
