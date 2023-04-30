package main

import "fmt"

// Variables globales
var mensaje string

func funcion1() {
	mensaje = "Hola desde funcion1"
	fmt.Println(mensaje)

}

func funcion2() {
	mensaje := "Hola desde funcion2"
	fmt.Println(mensaje)
}

func main() {
	mensaje = "Hola desde main"
	fmt.Println(mensaje)

	funcion1()
	funcion2()

	fmt.Println(mensaje)
}
