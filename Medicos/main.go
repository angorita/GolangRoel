package main

import (
	"Medicos/u"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//slice de balances
	var balances []u.Bal

	// Cargar contactos existentes desde el archivo
	err := u.LoadBalanceFromFile(&balances)
	if err != nil {
		fmt.Println("Error al cargar los balances:", err)
	}
	//crear instancia de bufio
	reader := bufio.NewReader(os.Stdin)
	for {
		//mostrar opciones del usuario

		fmt.Print("==== Balances Mensuales SSC ====\n",
			"1. Agregando un Balance\n",
			"2. Mostrar todos los Balances\n",
			"3. Salir\n",
			"Elige una opción: ")

		//leer la opcion del usuario
		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opcion:", err.Error())
			return
		}
		//manejar la opcion del usuario
		switch option {
		case 1:
			//ingresar y crear balance
			var b u.Bal

			fmt.Print("MesAnio: ")
			mesanio, _ := reader.ReadString('\n')
			b.MesAnio, _ = strconv.Atoi(strings.TrimSpace(mesanio))

			fmt.Print("Combo96: ")
			combo96, _ := reader.ReadString('\n')
			b.Combo96, _ = strconv.ParseFloat(strings.TrimSpace(combo96), 64)

			fmt.Print("HorasExtras: ")
			horasextras, _ := reader.ReadString('\n')
			b.HorasExtras, _ = strconv.ParseFloat(strings.TrimSpace(horasextras), 64)

			fmt.Print("DiferenciaCaja: ")
			diferenciacaja, _ := reader.ReadString('\n')
			b.DiferenciaCaja, _ = strconv.ParseFloat(strings.TrimSpace(diferenciacaja), 64)

			fmt.Print("DiferenciaAguinaldo: ")
			diferenciaaguinaldo, _ := reader.ReadString('\n')
			b.DiferenciaAguinaldo, _ = strconv.ParseFloat(strings.TrimSpace(diferenciaaguinaldo), 64)

			fmt.Print("DiferenciaBlanco: ")
			diferenciablanco, _ := reader.ReadString('\n')
			b.DiferenciaBlanco, _ = strconv.ParseFloat(strings.TrimSpace(diferenciablanco), 64)

			//agregar un balance a slice
			balances = append(balances, b)
			//guardar en archivo json
			if err := u.SaveBalanceToFile(balances); err != nil {
				fmt.Println("Error al guardar el balance", err)
			}
		case 2:
			//mostrar los balances ssc
			fmt.Println(`
		****************************************************`)
			for index, balance := range balances {
				fmt.Printf(`
		(%d). 
		----------------------------------------------------
		MesAnio: %d
		Combo96: %.2f
		HorasExtras: %.2f
		DiferenciaCaja: %.2f
		DiferenciaAguinaldo: %.2f
		DiferenciaBlanco: %.2f
		Total: %.2f
		----------------------------------------------------
					`,
					index+1, balance.MesAnio, balance.Combo96, balance.HorasExtras, balance.DiferenciaCaja, balance.DiferenciaAguinaldo, balance.DiferenciaBlanco, balance.CalcularTotal())
			}
			fmt.Println(`
		***************************************************** `)
		case 3:
			// Salir del programa
			fmt.Println("Saliendo de Programa")
			return
		default:
			fmt.Println("Opción no válida")
		}

	}

}
