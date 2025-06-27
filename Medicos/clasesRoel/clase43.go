package main

import (
	"errors"
	"fmt"
	"strconv"
)

// funcion que si divide por cero me larga un mensaje con errors.New("Aca va el tipo de error")
func Divide(div, endo int) (int, error) {
	if endo == 0 {
		return 0, errors.New("No se puede dividir por cero")
	}
	return div / endo, nil
}

func main() {
	num, e := Divide(10, 0)
	if e != nil {
		fmt.Println("Error: ", e)
		return
	}
	fmt.Println("Resultados", num)

	str := "234w"
	file, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("NUmero ", file)
}
