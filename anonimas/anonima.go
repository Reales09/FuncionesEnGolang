package main

import "fmt"

func main() {
	//Funcion anonima

	/*
		func() {
			fmt.Println("Hola desde la funciona anonima")
		}()
	*/

	myfunc := func(nombre string) string {
		return fmt.Sprintf("Hola %s desde la funcion anonima", nombre)

	}

	fmt.Println(myfunc("Reales"))

}
