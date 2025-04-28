package entities

import "fmt"

type Animal struct {
	Name string
	Age  int
}

func (animal Animal) Speak() {
	fmt.Println("Hello human")
}
