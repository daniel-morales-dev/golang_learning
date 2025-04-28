package entities

import "fmt"

type Zoo struct {
	animals []Animal
}

func NewZoo() *Zoo {
	return &Zoo{}
}

func (z *Zoo) AddAnimal(a Animal) {
	z.animals = append(z.animals, a)
}

func (z Zoo) ShowSounds() {
	if len(z.animals) == 0 {
		fmt.Println("No hay animales en el zoológico.")
		return
	}
	fmt.Println("Animales en el zoológico:")
	for _, animal := range z.animals {
		fmt.Println(animal.Speak())
	}
}
