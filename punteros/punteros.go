package main

import "fmt"

func modificarValor(cadena *string) {
	fmt.Printf("%p\n", cadena)
	*cadena = "Hola desde la funcion"
}

func main() {
	cadena := "Hola mundo de Go"
	fmt.Printf("%p\n", &cadena)
	fmt.Printf("Antes de la funcion: %s\n", cadena)

	modificarValor(&cadena) //copia

	fmt.Printf("Despues de la funcion: %s\n", cadena)

}
