package main

import (
	"fmt"
	"zoo-app/entities"
)

func main() {
	zoo := entities.NewZoo()
	fmt.Println("Zoo creado")

	Tiger := entities.NewAnimal("Tigre", "bengala", 2, "Miau!")

	fmt.Println("Animal creado:", Tiger.Name)

	fmt.Println("Agregando animal al Zoo...")

	zoo.AddAnimal(*Tiger)

	zoo.ShowSounds()
}
