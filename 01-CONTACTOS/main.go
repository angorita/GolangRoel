package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var contactos []Contacto
	err := LoadContactosFromFile(&contactos)
	if err != nil {
		fmt.Println("Error al cargar los contactos: ", err)
	}
	// crear instancia de fubio
	reader := bufio.NewReader(os.Stdin)
	for {
		//mostrar opciones al usuario
		fmt.Println("=====Gestor de contactos======", "1.Agregar un contacto\n", "2.Mostrar todos los contactos", "3.Salir\n")
	}
}
