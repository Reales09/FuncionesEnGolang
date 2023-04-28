package main

import "fmt"

func sumarReturnVarios(nombre string, numeros ...int) (string, int) {
	mensaje := fmt.Sprintf("La suma de %s es: ", nombre)
	fmt.Println(numeros)
	var total int
	for _, numero := range numeros {
		fmt.Println(numero)
		total += numero
	}
	return mensaje, total
}

func main() {
	mensaje, result := sumarReturnVarios("Reales", 1, 2, 3, 4, 5, 70)

	fmt.Println(mensaje, result)
}
