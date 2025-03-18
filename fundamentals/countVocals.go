package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var vocals = []string{"a", "e", "i", "o", "u"}
	var vocalsCounting = map[string]int{
		"a": 0,
		"e": 0,
		"i": 0,
		"o": 0,
		"u": 0,
	}

	reader := bufio.NewReader(os.Stdin) // Crea un lector de entrada
	//var count int

	var phrase string

	fmt.Println("Input a phrase:")
	phrase, _ = reader.ReadString('\n') // Captura la línea completa hasta "Enter"

	fmt.Println("Phrase:", phrase)

	for i := 0; i < len(phrase); i++ {
		value, ok := vocalsCounting[string(phrase[i])]
		if ok {
			vocalsCounting[string(phrase[i])] = value + 1
		}
	}

	fmt.Println("Counting vocals:", vocalsCounting)

	var count int
	for _, value := range vocalsCounting {
		count += value
	}

	fmt.Println("Total vocals:", count)

}
