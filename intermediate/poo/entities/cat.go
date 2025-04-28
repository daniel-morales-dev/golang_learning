package entities

import "fmt"

type Cat struct {
	Animal
	Weight float32
	Color  string
}

func (cat Cat) Speak() {
	fmt.Println("Miau!")
}
