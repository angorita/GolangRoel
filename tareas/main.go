package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	Nombre     string
	Desc       string
	Completado bool
}
type ListaTareas struct {
	Tareas []Tarea //slice que acumula las tareas
}

// recibe un puntero lista devuelve metodo agregartarea
func (l *ListaTareas) agregarTarea(t Tarea) {
	l.Tareas = append(l.Tareas, t)
}

// recibe un index con el cual buscamos en tareas y lo pasamos a true
func (l *ListaTareas) marcarCompletado(index int) {
	l.Tareas[index].Completado = true //cambiamos el false por true
}

// receptor es ListaTareas y recibimos la tarea que queremos editar...
func (l *ListaTareas) editarTarea(index int, t Tarea) {
	l.Tareas[index] = t
}
func (l *ListaTareas) eliminarTarea(index int) {
	l.Tareas = append(l.Tareas[:index], l.Tareas[index+1:]...)
}
func main() {
	//instancia
	lista := ListaTareas{} //objeto vacio de lista
	leer := bufio.NewReader(os.Stdin)
	var option int
	for {
		fmt.Println(`

		Seleccione una opcion :
			1)Agregar tarea
			2)Marcar tarea como completada
			3)Editar tarea
			4)Eliminar tarea
			5)Salir		
		
		`)
		fmt.Print(`
		Ingrese el numero de opcion: `)
		fmt.Scanln(&option)
		switch option {
		case 1:
			var t Tarea
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.Nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la descripcion de tarea: ")
			t.Desc, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea agregada : ")
		case 2:
			var index int
			fmt.Print("Ingrese el index de tarea completada: ")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("Tarea marcada como completada: ")
		case 3:
			var index int
			var t Tarea
			fmt.Print("Ingrese el index de la tarea a modificar: ")
			fmt.Scanln(&index)
			fmt.Print("Ingrese el nombre de la tarea : ")
			t.Nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la descripcion de la tarea : ")
			t.Desc, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
			fmt.Print("Tarea modificada")

		case 4:
			var index int
			fmt.Print("Ingrese el index de la tarea a Eliminar: ")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Print("Tarea eliminada")
		case 5:
			fmt.Print("Saliendo del programa")
			return
		default:
			fmt.Print("Opcion invalida")
		}

		fmt.Println("Lista de Tareas")
		fmt.Println(`====================================================================`)
		for i, t := range lista.Tareas {
			//recorro el objeto con Printf
			fmt.Printf(" %d.%s-%s -Completado-: %t\n", i, t.Nombre, t.Desc, t.Completado)
		}
		fmt.Println("====================================================================")
	}

}
