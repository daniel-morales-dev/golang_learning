package main

import (
	"animals/entities"
	"animals/interfaces"
	"fmt"
)

func main() {
	var animal []interfaces.Speaker

	dog := entities.Dog{Weight: 30.5, Animal: entities.Animal{Name: "Nesquik", Age: 2}}
	cat := entities.Cat{Weight: 30.9, Color: "Black", Animal: entities.Animal{Name: "Vaca", Age: 1}}

	animal = append(animal, dog)
	animal = append(animal, cat)

	fmt.Println("Animals", animal)
}
