package main

import "fmt"

type Student struct {
	Name          string
	Age           int
	Qualification float32
}

func (s *Student) Birthday() {
	s.Age += 1 // Modifica el struct original
}

func (s *Student) UpdateQualification(value float32) {
	s.Qualification = value
}

func main() {
	var Daniel = Student{Name: "Daniel", Age: 25}
	fmt.Printf("Nombre: %s | Edad: %d | Calificación: %.2f\n", Daniel.Name, Daniel.Age, Daniel.Qualification)

	fmt.Println("Dia de cumple!")
	Daniel.Birthday()

	fmt.Printf("Nombre: %s | Edad: %d | Calificación: %.2f\n", Daniel.Name, Daniel.Age, Daniel.Qualification)

	fmt.Println("Update qualification!")
	Daniel.UpdateQualification(3.8)
	fmt.Println("Daniel:", Daniel)
	fmt.Printf("Nombre: %s | Edad: %d | Calificación: %.2f\n", Daniel.Name, Daniel.Age, Daniel.Qualification)

}
