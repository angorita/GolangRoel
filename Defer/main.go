package main

import (
	"fmt"
	"os"
)

// "fmt"
// "os"
// PARA MOSTRAR COMO TRABAJA DEFER DEMORA HASTA QUE SE CIERRA O SI NO HAY ERROR
// usa recove() en vez de panic, continuando con la ejecucion...
func main() {
	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("Hola, Oscar Angarita"))
	if err != nil {
		fmt.Println(err)
		return
	}

	//	uso de panico

	divide(12, 1)
	divide(12, 1)
	divide(13, 0)
	divide(12, 1)
	divide(12, 3)
	divide(12, 4)

}
func divide(di, de int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	valiZero(de)
	fmt.Println(di / de)
}
func valiZero(de int) {
	if de == 0 {
		panic("Error no se puede dividir por 0")
	}
}
