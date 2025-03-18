package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	Name              string
	Price             float32
	AvailableQuantity int
}

var productList []Product
var cacheProduct = make(map[string]int)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var exit = 1

	for exit != 0 {
		fmt.Println("Que deseas hacer?\n1. Agregar producto.\n2. Buscar producto\n3. Calcular valor del inventario")
		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Println("Ingresa el nombre del producto")
			name, _ := reader.ReadString('\n')
			name = normalizeProductName(name)

			fmt.Println("Ingresa su cantidad")
			var availableQuantity int
			fmt.Scanln(&availableQuantity)

			fmt.Println("Ingresa su precio")
			var price float32
			fmt.Scanln(&price)

			newProduct := Product{Name: name, AvailableQuantity: availableQuantity, Price: price}
			addProduct(newProduct)
		case 2:
			fmt.Println("Ingresa el nombre del producto")
			name, _ := reader.ReadString('\n')
			name = normalizeProductName(name)
			searchProduct(name)
		case 3:
			fmt.Println("Calculando valor del inventario...")
			var total = calculateInventory()
			fmt.Printf("El valor del inventario es %g\n", total)
		default:
			fmt.Println("Opción fuera de rango")
		}
		fmt.Println("===========================")
		fmt.Println("Deseas salir?\n0. Si\n1. No")
		fmt.Scanln(&exit)
	}
}

func addProduct(p Product) {
	productList = append(productList, p)
	cacheProduct[p.Name] = len(productList) - 1
}

func searchProduct(name string) {
	product := Product{}

	index, ok := cacheProduct[name]
	if ok {
		fmt.Println(printProduct(productList[index]))
	} else {
		fmt.Println("Producto no encontrado")
	}
}

func printProduct(product Product) string {
	return "Producto: " + product.Name + " | Cantidad: " + strconv.Itoa(product.AvailableQuantity) + " | Precio: " + strconv.FormatFloat(float64(product.Price), 'f', 2, 32) + "\n"
}

func calculateInventory() float32 {
	var total float32
	for _, product := range productList {
		total += product.Price * float32(product.AvailableQuantity)
	}

	return total
}

func normalizeProductName(name string) string {
	return strings.ToLower(strings.TrimSpace(name))
}
