package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Estructura de contactos
type Contact struct {
	Descripcion string  `json:"descripcion"`
	Precio      float64 `json:"precio"`
	Cantidad    int     `json:"cantidad"`
	Fecha       string  `json:"fecha"`
	Dolar       float64 `json:"dolar"`
	Id          int     `json:"id"`
	Bhabilitado bool    `json:"bhabilitado"`
}

func main() {
	// Slice de contactos
	var contacts []Contact

	// Cargar contactos existentes desde el archivo
	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar los contactos:", err)
	}

	// Crear instancia de bufio
	reader := bufio.NewReader(os.Stdin)

	for {
		// Mostrar opciones al usuario
		fmt.Print("==== Materiales Loft Oscar ====\n",
			"1. Agregar un Material\n",
			"2. Mostrar todos los Materiales\n",
			"3. Salir\n",
			"Elige una opción: ")
		// Leer la opción del usuario
		var option int
		_, err := fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opción:", err)
			return
		}

		// Manejar la opción del usuario
		switch option {
		case 1:
			// Ingresar y crear contacto
			var c Contact

			fmt.Print("Descripcion: ")
			desc, _ := reader.ReadString('\n')
			c.Descripcion = strings.TrimSpace(desc)

			fmt.Print("Precio: ")
			precioStr, _ := reader.ReadString('\n')
			c.Precio, _ = strconv.ParseFloat(strings.TrimSpace(precioStr), 64)
			fmt.Print("Cantidad: ")
			cantStr, _ := reader.ReadString('\n')
			c.Cantidad, _ = strconv.Atoi(strings.TrimSpace(cantStr))

			fmt.Print("Fecha: ")
			fecha, _ := reader.ReadString('\n')
			c.Fecha = strings.TrimSpace(fecha)

			fmt.Print("Dolar: ")
			dolarStr, _ := reader.ReadString('\n')
			c.Dolar, _ = strconv.ParseFloat(strings.TrimSpace(dolarStr), 64)

			fmt.Print("Id: ")
			idStr, _ := reader.ReadString('\n')
			c.Id, _ = strconv.Atoi(strings.TrimSpace(idStr))

			fmt.Print("Bhabilitado (true/false): ")
			habilitadoStr, _ := reader.ReadString('\n')
			c.Bhabilitado, _ = strconv.ParseBool(strings.TrimSpace(habilitadoStr))
			// Agregar un contacto a Slice
			contacts = append(contacts, c)
			// Guardar en un archivo json
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error al guardar el contacto:", err)
			}

		case 2:
			// Mostrar todos los contactos
			fmt.Println(`
		====================================================`)
			for index, contact := range contacts {
				fmt.Printf(`
		%d. 
		Descripción: %s
		Precio: %.2f
		Cantidad: %d
		Fecha: %s
		Dólar: %.2f
		Id: %d
		Habilitado: %t
		----------------------------------------------------
					`,
					index+1, contact.Descripcion, contact.Precio, contact.Cantidad,
					contact.Fecha, contact.Dolar, contact.Id, contact.Bhabilitado)
			}
			fmt.Println(`
		====================================================
		`)
		case 3:
			// Salir del programa
			return
		default:
			fmt.Println("Opción no válida")
		}

	}
}

// Guarda contatos en un archivo json
func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}
	return nil
}

// Carga contactos desde un archivo json
func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}

	return nil
}
