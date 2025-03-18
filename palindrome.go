package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input a phrase:")

	phrase, _ := reader.ReadString('\n')

	phrase = strings.ReplaceAll(strings.ToLower(strings.TrimSpace(phrase)), " ", "")

	fmt.Println("Phrase:", phrase)

	realPhrase := ""

	for i := len(phrase)-1;i>=0;i--{
		realPhrase += string(phrase[i])
	}

	fmt.Println("Real phrase:", realPhrase)

	if realPhrase == phrase {
		fmt.Println("Is a palindrome")
	} else {
		fmt.Println("Is not a palindrome")
	}
}
