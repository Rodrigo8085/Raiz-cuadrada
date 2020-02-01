package main

import (
	"fmt"
	"math"
)

func raizCuadrada(numero float64) (float64, error) {
	if numero < 0 {
		return 0, fmt.Errorf("No se puede obtener la raiz cuadrada de un numero negativo")
	}
	return math.Sqrt(numero), nil
}

func main() {
	raiz, err := raizCuadrada(-9.3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("%0.3f", raiz)
	}
}
