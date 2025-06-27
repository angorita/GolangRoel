package main

import (
	"fmt"
	"web2/db"
	"web2/models"
)

func main() {
	db.Connect()
	// material := models.CreateMaterial("Yaves", 122.33, 1, "2025/12/1", 120, true)
	// fmt.Println(material)
	// fmt.Println(db.ExistsTable("material"))
	// db.CreateTable(models.MaterialSchema, "material")
	// db.TruncateTable("material"
	// db.Ping()
	// materiales := models.ListarMateriales()
	// fmt.Println(materiales)
	// materiales2 := models.GetMateriales(1)
	// fmt.Println(materiales2)
	//busca el material 3, luego lo imprime, luego lo modifica y lo vuelve a imprimr despues cuando guardamos
	// material := models.GetMateriales(3)
	// fmt.Println(material)
	// material.Descripcion = "Hierro Doblado con Maestria"
	// material.Dolar = 1000
	// material.Bhabilitado = false
	//ahora llamo a save() para guardar, obvio no?
	// material.Save()
	//como borrar ? busco primero despues borro y despues
	material := models.GetMateriales(1)
	fmt.Println("Muestro el registro a borrar", material)
	material.Delete()
	fmt.Println(models.ListarMateriales())
	db.TruncateTable("material")
	db.Close()
}
