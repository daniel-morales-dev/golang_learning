package main

import "fmt"

// Definimos una interface
type Notifier interface {
	Notify(message string)
}

// Struct que implementa la interface
type Email struct {
	Email string
}

func (e Email) Notify(message string) {
	fmt.Printf("Enviando email a %s: %s\n", e.Email, message)
}

// Otro struct que implementa la misma interface
type SMS struct {
	Number string
}

func (s SMS) Notify(message string) {
	fmt.Printf("Enviando SMS a %s: %s\n", s.Number, message)
}

// Función que recibe cualquier tipo que implemente la interface
func SendNotification(n Notifier, message string) {
	n.Notify(message)
}

func main() {
	email := Email{Email: "daniel@example.com"}
	sms := SMS{Number: "+12345678"}

	SendNotification(email, "¡Hola por email!")
	SendNotification(sms, "¡Hola por SMS!")
}
