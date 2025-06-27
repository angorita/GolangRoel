package main

import (
	"fmt"
	"math/rand"
)

func main() {
	Jugar()
}
func Jugar() {
	numAleatorio := rand.Intn(3)
	var numeroIngresado int
	var intentos int
	const maxIntentos = 3

	for intentos < maxIntentos {
		intentos++
		fmt.Printf(`Segui ingresando numeros hasta que el numero de intentos %d termine :`, maxIntentos-intentos+1)
		fmt.Scanln(&numeroIngresado)
		if numeroIngresado == numAleatorio {
			fmt.Println("Felicitaciones la pegaste ... pedazo de suertudo!!!")
			JugarNuevamente()
			return
		} else if numeroIngresado < numAleatorio {
			fmt.Println("Pone un numero mayor")
		} else if numeroIngresado > numAleatorio {
			fmt.Println("Pone un numero menor!!!")
		}
	}
	fmt.Println(`Se acabaron los intentos el numero era : `, numAleatorio)
	JugarNuevamente()
}
func JugarNuevamente() {
	var eleccion string
	fmt.Println("Quieres jugar nuevamente? S/N ?")
	fmt.Scanln(&eleccion)
	switch eleccion {
	case "s":
		Jugar()
	case "n":
		fmt.Println("Bueno, Chauuuu!!!")
	default:
		fmt.Println("Seleccion invalida.!!!")
		JugarNuevamente()
	}
}
