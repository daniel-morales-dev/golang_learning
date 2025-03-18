package main

import (
	"fmt"
	"strings"
)

func main() {
	var people = make(map[string]int)
	for i := 0; i < 3; i++ {
		fmt.Println("Ingresa un nombre:")
		var name string
		fmt.Scanln(&name)
		name = strings.ToUpper(name)
		fmt.Println("Ingresa su edad:")
		var age int
		fmt.Scanln(&age)
		people[name] = age
	}

	fmt.Println("Personas: ", people)

	fmt.Println("Ingrese un nombre:")
	var name string
	fmt.Scan(&name)
	name = strings.ToUpper(name)

	age, ok := people[name]
	if ok {
		fmt.Printf("La persona %v tiene una edad de %v", name, age)
	} else {
		fmt.Println("Esa persona no esta en nuestro diccionario")
	}
}
