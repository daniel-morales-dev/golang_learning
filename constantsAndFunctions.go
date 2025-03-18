package main

import (
	"fmt"
	"strconv"
)

func main() {
	const fullName string = "Daniel Morales"
	const age int = 25

	fmt.Println(sayHello(fullName, age))
}

func sayHello(name string, age int) string {
	return "Hello, " + name + " you are " + strconv.Itoa(age) + " years old"
}
