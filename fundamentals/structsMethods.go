package main

import "fmt"

type Student struct {
	Name          string
	Age           int
	Qualification float32
}

// Lo importante es (s Student), que es el receiver. Le indica a Go que este método pertenece a Student
// Esto es un campo por valor, por ende no modificara el struct original
func (s Student) SayHello() {
	fmt.Printf("Hello, I am %s and I am %d years old\n", s.Name, s.Age)
}

func (s Student) IsApproved() bool {
	return s.Qualification >= 3.0
}

func main() {
	var Daniel = Student{Name: "Daniel", Age: 25, Qualification: 5.0}

	Daniel.SayHello()
	fmt.Printf("%s Aprobó? %t\n", Daniel.Name, Daniel.IsApproved())
}
