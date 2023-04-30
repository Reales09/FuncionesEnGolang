package main

import "fmt"

func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("No se puede dividir por cero")
	} else {
		return dividendo / divisor, nil
	}

}

func main() {
	resultado, err := division(4, 2)
	if err == nil {
		fmt.Println("El resultado es:", resultado)
	} else {
		fmt.Println(err)
	}

}
