package main

import "fmt"

func sumar(numeros ...int) int {
	fmt.Println(numeros)
	var total int
	for _, numero := range numeros {
		fmt.Println(numero)
		total += numero
	}
	return total
}

func main() {
	result := sumar(1, 2, 3, 4, 5, 70)

	fmt.Println(result)
}
