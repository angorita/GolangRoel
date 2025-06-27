package main

import (
	"Errores/funciones"
	"fmt"
)

func main() {
	//llamar a la funcion divide
	r, e := funciones.Divide(12, 1)
	if e != nil {
		fmt.Println("Error: ", e)
		return
	}
	defer fmt.Println("El resultado es : ", r, " Nanogramos")
}
