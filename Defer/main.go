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

	file, err := os.Create("txtCreadoPorOAA.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte("Hola, Oscar Angarita, como andas chaval ?"))
	if err != nil {
		fmt.Println(err)
		return
	}

	//	uso de panico

	divide(12, 1)
	divide(13, 0)
	divide(12, 3)
	divide(12, 4)
	/* 	log.SetPrefix("en main ocurren cosas main():")
	   	log.Print("Mensaje")
	   	log.Panic("Horror!!!") */

}

/*
La palabra recover va con defer para capturar y manejar panic
en una version diferida permite seguir al programa.
se llama a recover dentro de una funcion defer
*/
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
		panic("👋👋👋valiZero👋👋👋")
	}
}
