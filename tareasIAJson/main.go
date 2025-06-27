package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	opcionAgregar   = 1
	opcionCompletar = 2
	opcionEditar    = 3
	opcionEliminar  = 4
	opcionSalir     = 5
	nombreArchivo   = "tareas.json" // Nombre del archivo JSON para guardar las tareas
)

type Tarea struct {
	Nombre     string
	Desc       string
	Completado bool
}

type ListaTareas struct {
	Tareas []Tarea
}

func (l *ListaTareas) agregarTarea(t Tarea) {
	l.Tareas = append(l.Tareas, t)
}

func (l *ListaTareas) marcarCompletado(index int) error {
	if index < 0 || index >= len(l.Tareas) {
		return fmt.Errorf("índice de tarea inválido: %d", index)
	}
	l.Tareas[index].Completado = true
	return nil
}

func (l *ListaTareas) editarTarea(index int, t Tarea) error {
	if index < 0 || index >= len(l.Tareas) {
		return fmt.Errorf("índice de tarea inválido: %d", index)
	}
	l.Tareas[index] = t
	return nil
}

func (l *ListaTareas) eliminarTarea(index int) error {
	if index < 0 || index >= len(l.Tareas) {
		return fmt.Errorf("índice de tarea inválido: %d", index)
	}
	l.Tareas = append(l.Tareas[:index], l.Tareas[index+1:]...)
	return nil
}

func (l *ListaTareas) imprimirListaTareas() {
	fmt.Println("\nLista de Tareas")
	fmt.Println("====================================================================")
	if len(l.Tareas) == 0 {
		fmt.Println("No hay tareas en la lista.")
	} else {
		for i, t := range l.Tareas {
			estado := "Pendiente"
			if t.Completado {
				estado = "Completada"
			}
			fmt.Printf("%d. %s - %s (%s)\n", i+1, strings.TrimSpace(t.Nombre), strings.TrimSpace(t.Desc), estado)
		}
	}
	fmt.Println("====================================================================")
}

func guardarTareas(l *ListaTareas) error {
	data, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return fmt.Errorf("error al convertir las tareas a JSON: %w", err)
	}
	err = ioutil.WriteFile(nombreArchivo, data, 0644)
	if err != nil {
		return fmt.Errorf("error al guardar las tareas en el archivo '%s': %w", nombreArchivo, err)
	}
	fmt.Println("Tareas guardadas en:", nombreArchivo)
	return nil
}

func cargarTareas() ListaTareas {
	data, err := ioutil.ReadFile(nombreArchivo)
	if err != nil {
		fmt.Println("No se encontró el archivo de tareas o hubo un error al leerlo. Iniciando con una lista vacía.")
		return ListaTareas{}
	}
	var lista ListaTareas
	err = json.Unmarshal(data, &lista)
	if err != nil {
		fmt.Println("Error al cargar las tareas desde el archivo. Iniciando con una lista vacía.")
		return ListaTareas{}
	}
	fmt.Println("Tareas cargadas desde:", nombreArchivo)
	return lista
}

func main() {
	lista := cargarTareas()
	leer := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf(`
Seleccione una opción:
  %d) Agregar tarea
  %d) Marcar tarea como completada
  %d) Editar tarea
  %d) Eliminar tarea
  %d) Salir
Ingrese el número de opción: `, opcionAgregar, opcionCompletar, opcionEditar, opcionEliminar, opcionSalir)

		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error al leer la entrada. Intente nuevamente.")
			continue
		}

		option, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Opción inválida. Ingrese un número.")
			continue
		}

		switch option {
		case opcionAgregar:
			var t Tarea
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.Nombre, _ = leer.ReadString('\n')
			t.Nombre = strings.TrimSpace(t.Nombre)
			fmt.Print("Ingrese la descripción de la tarea: ")
			t.Desc, _ = leer.ReadString('\n')
			t.Desc = strings.TrimSpace(t.Desc)
			lista.agregarTarea(t)
			fmt.Println("Tarea agregada.")

		case opcionCompletar:
			if len(lista.Tareas) == 0 {
				fmt.Println("No hay tareas para marcar como completadas.")
				continue
			}
			fmt.Print("Ingrese el número de la tarea para marcar como completada: ")
			var indexInput string
			_, err := fmt.Scanln(&indexInput)
			if err != nil {
				fmt.Println("Error al leer la entrada. Intente nuevamente.")
				continue
			}
			index, err := strconv.Atoi(indexInput)
			if err != nil {
				fmt.Println("Índice inválido. Ingrese un número.")
				continue
			}
			err = lista.marcarCompletado(index - 1) // Restamos 1 porque los índices en la lista son base 0
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Tarea marcada como completada.")
			}

		case opcionEditar:
			if len(lista.Tareas) == 0 {
				fmt.Println("No hay tareas para editar.")
				continue
			}
			fmt.Print("Ingrese el número de la tarea a editar: ")
			var indexInput string
			_, err := fmt.Scanln(&indexInput)
			if err != nil {
				fmt.Println("Error al leer la entrada. Intente nuevamente.")
				continue
			}
			index, err := strconv.Atoi(indexInput)
			if err != nil {
				fmt.Println("Índice inválido. Ingrese un número.")
				continue
			}
			if index < 1 || index > len(lista.Tareas) {
				fmt.Printf("Índice de tarea inválido. Debe estar entre 1 y %d.\n", len(lista.Tareas))
				continue
			}
			var t Tarea
			fmt.Print("Ingrese el nuevo nombre de la tarea: ")
			t.Nombre, _ = leer.ReadString('\n')
			t.Nombre = strings.TrimSpace(t.Nombre)
			fmt.Print("Ingrese la nueva descripción de la tarea: ")
			t.Desc, _ = leer.ReadString('\n')
			t.Desc = strings.TrimSpace(t.Desc)
			err = lista.editarTarea(index-1, t)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Tarea modificada.")
			}

		case opcionEliminar:
			if len(lista.Tareas) == 0 {
				fmt.Println("No hay tareas para eliminar.")
				continue
			}
			fmt.Print("Ingrese el número de la tarea a eliminar: ")
			var indexInput string
			_, err := fmt.Scanln(&indexInput)
			if err != nil {
				fmt.Println("Error al leer la entrada. Intente nuevamente.")
				continue
			}
			index, err := strconv.Atoi(indexInput)
			if err != nil {
				fmt.Println("Índice inválido. Ingrese un número.")
				continue
			}
			err = lista.eliminarTarea(index - 1) // Restamos 1 porque los índices en la lista son base 0
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Tarea eliminada.")
			}

		case opcionSalir:
			err := guardarTareas(&lista)
			if err != nil {
				fmt.Println("Error al guardar las tareas:", err)
			}
			fmt.Println("Saliendo del programa.")
			return

		default:
			fmt.Println("Opción inválida.")
		}

		lista.imprimirListaTareas()
	}
}
