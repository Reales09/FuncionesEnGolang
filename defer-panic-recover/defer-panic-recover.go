package main

import (
	"fmt"
	"os"
)

func main() {

	// Nos devuelve una variable de tipo file y un error

	if file, error := os.Open("C:/Users/DESARROLLADOR4/go/src/curso_golang_udemy/funciones/defer-panic-recover/hol.txt"); error != nil {
		panic("No fue posible obtener el archivo!!")
	} else {
		defer func() {
			fmt.Println("Cerrando el archivo!!")
			file.Close()
		}()
		contenido := make([]byte, 254)
		long, _ := file.Read(contenido)

		contenidoArchivo := string(contenido[:long])
		fmt.Println(contenidoArchivo)

	}

}
