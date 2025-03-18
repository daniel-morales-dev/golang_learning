package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	wordsCounting := map[string]int{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input a phrase:")

	phrase, _ := reader.ReadString('\n')

	words := strings.Fields(strings.ToLower(strings.TrimSpace(phrase)))

	for _, word := range words {
		wordsCounting[word]++
	}

	fmt.Println("Words counting:", wordsCounting)

}
