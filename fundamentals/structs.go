package main

import "fmt"

type Student struct {
	Name          string
	Age           int
	Qualification float32
}

func main() {
	var exit = 1
	var students []Student
	for exit != 0 {
		fmt.Println("Ingresa Nombre, Edad y Calificación, separado por espacios")
		var name string
		var age int
		var qualification float32
		fmt.Scanln(&name, &age, &qualification)

		students = append(students, Student{Name: name, Age: age, Qualification: qualification})

		fmt.Println("Deseas salir? (0 para salir, 1 para continuar)")
		fmt.Scanln(&exit)
	}
	var prom float32
	var sumQualification float32
	for i, student := range students {
		fmt.Printf("Estudiante %d: Nombre: %s | Edad: %d | Calificación: %.2f\n", i+1, student.Name, student.Age, student.Qualification)
	}

	for _, student := range students {
		sumQualification += student.Qualification
	}

	prom = sumQualification / float32(len(students))

	fmt.Println("El promedio de calificaciones fue:", prom)

}
