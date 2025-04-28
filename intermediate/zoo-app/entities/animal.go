package entities

import "fmt"

type Animal struct {
	Name  string
	Breed string
	Age   int
	Sound string
}

func NewAnimal(name string, breed string, age int, sound string) *Animal {
	return &Animal{Name: name, Breed: breed, Age: age, Sound: sound}
}

func (a Animal) Speak() string {
	return fmt.Sprintf("El %s hace: %s", a.Name, a.Sound)
}
