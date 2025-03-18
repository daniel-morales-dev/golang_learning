package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// En Go, las strings son secuencias de bytes (UTF-8), no caracteres individuales.
// - `byte` (alias de uint8) representa un solo byte y funciona bien para ASCII.
// - `rune` (alias de int32) representa un carácter Unicode, útil para tildes, emojis, etc.
// - Convertir `string` a `rune` permite recorrer correctamente caracteres especiales.
func main() {
	vocalsCounting := map[string]int{
		"a": 0, "e": 0, "i": 0, "o": 0, "u": 0,
	}
	totalVocals := 0

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input a phrase:")

	phrase, _ := reader.ReadString('\n')                // Captura la línea completa
	phrase = strings.ToLower(strings.TrimSpace(phrase)) // Normaliza la entrada

	fmt.Println("Phrase:", phrase)

	// Contar vocales
	for _, char := range phrase {
		charStr := string(char)
		if _, exists := vocalsCounting[charStr]; exists {
			vocalsCounting[charStr]++
			totalVocals++
		}
	}

	// Imprimir conteo de vocales
	fmt.Println("Counting vocals:", vocalsCounting)

	fmt.Println("Total vocals:", totalVocals)
}
