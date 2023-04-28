package main

import "fmt"

//Closure
func repeat(n int) func(cadena string) string {
	return func(cadena string) string {
		var result string
		for i := 0; i < n; i++ {
			result += cadena
		}
		return result
	}

}

func main() {

	closure := repeat(5)
	fmt.Println(closure("Hola"))
	//Funcion anonima1

	/*
			func() {
				fmt.Println("Hola desde la funciona anonima")
			}()


		myfunc := func(nombre string) string {
			return fmt.Sprintf("Hola %s desde la funcion anonima", nombre)

		}

		fmt.Println(myfunc("Reales"))
	*/

}
