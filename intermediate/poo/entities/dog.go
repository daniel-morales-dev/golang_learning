package entities

import "fmt"

type Dog struct {
	Animal
	Weight float32
}

func (dog Dog) Speak() {
	fmt.Println("WOOF WOOF")
}
