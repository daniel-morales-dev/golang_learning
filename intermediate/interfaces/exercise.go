package main

import "fmt"

type Printable interface {
	Print() string
}

type Book struct {
	Title  string
	Author string
}

func (book Book) Print() string {
	return fmt.Sprintf("Libro: %s\nAutor: %s\n", book.Title, book.Author)
}

type Magazine struct {
	Name   string
	Number int
}

func (magazine Magazine) Print() string {
	return fmt.Sprintf("Revista: %s\nNumero: %d\n", magazine.Name, magazine.Number)
}

type Newspaper struct {
	Name string
	Date string
}

func (newspaper Newspaper) Print() string {
	return fmt.Sprintf("Periódico: %s\nFecha: %s\n", newspaper.Name, newspaper.Date)
}

func main() {
	var printableList []Printable
	book := Book{Title: "Dune", Author: "Alguien"}
	magazine := Magazine{Name: "Semana", Number: 1}
	newspaper := Newspaper{Name: "El País", Date: "2025-03-18"}

	printableList = append(printableList, book)
	printableList = append(printableList, magazine)
	printableList = append(printableList, newspaper)

	for _, printable := range printableList {
		fmt.Println(printable.Print())

	}

}
