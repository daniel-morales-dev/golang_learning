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

	words := strings.Fields(phrase)

	fmt.Println("Words:", words)
	fmt.Println("Total words:", len(words))
}
