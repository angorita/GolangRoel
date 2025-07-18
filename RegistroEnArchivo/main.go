package main

import (
	"log"
	"os"
)

/* se crea un log archivo que acumula los errores y los va guardando, lo crea lo agrega y lo pone como solo lectura, 0644 es un valor octal para los permisos de usuario, puede escribirlo solo el propietario*/
func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Print("Esto es un error")
}
